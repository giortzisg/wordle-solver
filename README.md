# Wordle Solver

## What does wordle-solver do? 

Wordle-solver is a cli tool to help you with the daily wordle.

## How does wordle-solver work?

The algorithm that the wordle solver uses is based on [this 3Blue1Brown video][3b1b] and it uses information theory. Essentially we calculate a given entropy for every word in our corpus that gives us the best probability of finding the correct answer. That means that words that eliminate many results are favorable and thus we recommend them. We also use a frequency map `frequency_map.json` that gives us the frequency of every word in our corpus used in the english vocabulary. This way we tend to recommend words that are known, opposed to some unknown words that generally have a lower probability of being the correct answer. See the YT video for reference of how the frequency data were collected.


[3b1b]: https://www.youtube.com/watch?v=v68zYyaEmEA

## Performance

Our wordle-solver is specifically designed to target the original wordle game (however, you can also change it buy changing ```
{
  "hardMode": "true"
}
``` on the `config.json` file to include all possible 5 letter words). The algorithm solves correctly all given words in 6 tries 100% of the time.

The average performance with the word `alone` as the first guess is 3.37 tries (check the benchmark file for details).

## Features

- The algorithm uses go routines to calculate the probability of the words faster. For the word `alone` and with `1 thread` (on the normal implementation) it takes approximately 1 minute and 10 seconds to solve all the possible words, while with `6 threads` is takes approximately 20 seconds, which is a significant improvement.
- Ability to play with all possible 5 letter words by changing the config values.