import json
import math
import os 
from pathlib import Path

SCRIPT_DIR = os.path.dirname(os.path.abspath(__file__))
ROOT_DIR = Path(SCRIPT_DIR).parent.absolute()

def sigmoid(pos):
    return 1 / (1 + math.exp(-pos))

# create a probability map for all words depending on their frequency for faster runtime computation
# the frequency_map.json file has the used frequency of the specified words in the english vocabulary
with open(str(ROOT_DIR) + '/config/frequency_map.json') as json_file:
    data = json.load(json_file)
    sorted = list(sorted(data.items(), key=lambda item: item[1]))
    word_possibility = {}
    for x in range(len(sorted)):
        word_possibility[sorted[x][0]] = sigmoid((x*12-117800)/2203)
 
    with open(str(ROOT_DIR) + '/config/probability_map.json', 'w', encoding='utf-8') as f:
        json.dump(word_possibility, f, ensure_ascii=False, indent=4)
