

import json



jsonScriptureFilePath = "../verses.json"
print(jsonScriptureFilePath)

def getVerses(filePath):
    with open(filePath, "r") as json_file: 
        verseList = json.load(json_file) # list of obj(humanName, url, content)
        # get list of genesis verses
        genesisVerses = list(filter(lambda v: "Genesis" in v["humanName"], verseList))
        # get list of exodus verses
        exodusVerses = list(filter(lambda v: "Exodus" in v["humanName"], verseList))
    return verseList, genesisVerses, exodusVerses



########################################################
## Load the data
########################################################

# Load in the data
allVerses, genesisVerses, exodusVerses = getVerses(jsonScriptureFilePath)
# print(genesisVerses[:5])
# print(exodusVerses[:5])

# genesisVerses = genesisVerses + exodusVerses # HACK - use both books now
genesisVerses = allVerses

########################################################
## Clean the data
########################################################

# ["“","‘","’","”","¶","1","2","3","4","5","6","7","8","9","0","\"","'",",",".","(",")",":","-"]
def removeCharactersAndPunctuation(s):
    for charStr in ["“","‘","’","”","¶","1","2","3","4","5","6","7","8","9","0","\"","'",",",".","(",")", ":", "-"]:
        s = s.replace(charStr, "")
    return s

def cleanVerses(verseArray):
    for verse in verseArray:
        content = verse["content"]
        content = removeCharactersAndPunctuation(content)
        content = content.strip()
        content = content.lower()
        verse["content"] = content
    return verseArray

genesisVerses = cleanVerses(genesisVerses)

########################################################
## Tokenize the datasets
########################################################

# Tokenize the datasets
# for collectionList in List(genesisVerses, exodusVerses):
#     for verseList in collectionList:
for verse in genesisVerses:
    verseContent = verse["content"]
    tokenList = verseContent.split()
    verse["tokens"] = tokenList

print("########################## Genesis Verses cleaned up ##############################")
print(genesisVerses[:2])


for verse in exodusVerses:
    verseContent = verse["content"]
    tokenList = verseContent.split()
    verse["tokens"] = tokenList

print(exodusVerses[:2])


########################################################
## Get it into a lines matrix ready for 
########################################################
print("--------------- Arranging into a Matrix ---------------")

inputMatrix = []
for verse in genesisVerses:
    tokenList = verse["tokens"]
    inputMatrix.append(tokenList)
print(inputMatrix[:10])

print("--------------- More Readable version ---------------")

for line in inputMatrix[:10]:
    print(line)



########################################################
## Glove NLP Model
########################################################

# importing the glove library
from glove import Glove, Corpus
# from Corpus 
# from fastai.basics import *

numComponents = 5

# creating a corpus object
corpus = Corpus() 

#training the corpus to generate the co occurence matrix which is used in GloVe
lines = inputMatrix
corpus.fit(lines)

#creating a Glove object which will use the matrix created in the above lines to create embeddings
#We can set the learning rate as it uses Gradient Descent and number of components
glove = Glove(no_components=numComponents)
 
glove.fit(corpus.matrix, no_threads=4, verbose=True)
glove.add_dictionary(corpus.dictionary)
# glove.save('glove.model')


####################################beast####################
## Create Embeddings for every LINE
########################################################

# get an embedding in a list that has the same indexes as the inputMatrix

# for every line
    # add up the embedding vectors for every word in the line
    # save those... 

import numpy as np

verseEmbeddingsList = []

# for 
for lineWordList in inputMatrix:
    lineEmbedding = np.zeros(numComponents)
    for word in lineWordList:
        # print(word)
        wordEmbedding = glove.word_vectors[glove.dictionary[word]]
        # print(wordEmbedding)
        lineEmbedding += wordEmbedding

    verseEmbeddingsList.append(lineEmbedding)

print(verseEmbeddingsList[:2])


########################################################
## SEARCH -- i.e. Sort based on cosine similarity 
########################################################

# CONVERT THIS TO A NUMPY ARRAY
verseEmbeddingsList = np.array(verseEmbeddingsList)

search = "praised"

# Sort the verse embeddings based on the cosine similarity of the search

searchEmbedding = glove.word_vectors[glove.dictionary[search]]

# comnpare searchEmbedding with embedding of every verse (cos similarity)

## get cosine similarity from every verse between that verse and my search term

from scipy import spatial

similarities = []
for verseEmbedding in verseEmbeddingsList:
    # https://stackoverflow.com/questions/18424228/cosine-similarity-between-2-number-lists
    cosSimilarity = 1.0 - spatial.distance.cosine(verseEmbedding, searchEmbedding)
    similarities.append(cosSimilarity)

print("########################################################")
print(similarities[10:14])


### getting indexes ??? 

indexes = np.argsort(-np.array(similarities))

for idx in indexes[:10]:
    print("################")
    print(idx)
    # print("Search embedding: ", searchEmbedding)
    # print("Line Embedding: ", verseEmbeddingsList[idx])
    print("Similarity: ", similarities[idx])
    print("VERSE : ", lines[idx])
    print("VERSE NAME : ", genesisVerses[idx]["humanName"])




########################################################
## Check synonyms
########################################################

# compare the search embedding with all the other words in the dictionary

synonymSearch = "conceived"
synonymSearchEmbedding = glove.word_vectors[glove.dictionary[synonymSearch]]



otherWordSimilarities = []
for word in glove.dictionary:
    idxToWordVectors = glove.dictionary[word]
    otherWordEmbedding = glove.word_vectors[idxToWordVectors]
    cosSimilarity = 1.0 - spatial.distance.cosine(synonymSearchEmbedding, otherWordEmbedding)
    otherWordSimilarities.append(cosSimilarity)


indexes = np.argsort(-np.array(otherWordSimilarities))

for index in indexes[:10]:
    print("Brian's words", glove.inverse_dictionary[index])

# for idx in indexes[:10]:
#     print("################")
#     print(idx)
#     # print("Search embedding: ", searchEmbedding)
#     # print("Line Embedding: ", verseEmbeddingsList[idx])
#     print("Similarity: ", otherWordSimilarities[idx])
#     # print("Word: ", )
#     # print("VERSE : ", lines[idx])
#     # print("VERSE NAME : ", genesisVerses[idx]["humanName"])







########################################################
## SEARCH -- i.e. Sort based on cosine similarity 
########################################################









