import json
import math


def sigmoid(pos):
    return 1 / (1 + math.exp(-pos))

with open('internal/entities/frequency_map.json') as json_file:
    data = json.load(json_file)
    sorted = list(sorted(data.items(), key=lambda item: item[1]))
    word_possibility = {}
    for x in range(len(sorted)):
        word_possibility[sorted[x][0]] = sigmoid((x*12-117800)/2203)
 
    with open('internal/entities/probability_map.json', 'w', encoding='utf-8') as f:
        json.dump(word_possibility, f, ensure_ascii=False, indent=4)
