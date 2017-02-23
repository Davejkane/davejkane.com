Goals vs Yellow Cards. Any Correlation?
A look at whether there is any correlation between yellow cards and goals
2016-11-30
## Hypothesis
So now that I have a nifty little database of statistics for the 2015/16 season of the English Premier League, it's to time to look at some more stats, and this week I'm going to look at whether there is any correlation between the number of yellow cards in a game and the number of goals in a game. We've got, as always, 380 data points, from all the games in a season.

Here's my hypothesis: The more yellow cards there are in a game, the more goals there will be. On average. I'm expecting a pretty weak correlation. My reasoning is, games with lots of yellow cards, are more likely to have red cards, and therefore lop-sided games leading to more goals. They're also more likely to have penalties. Again, more goals. They're maybe also higher paced and more action packed. I dunno, this is just my intuition. Let's have a look and see how full of shit I am.

## Goals
First lets have a quick look at the goals landscape. Some key figures:

* Average goals per game: 2.7
* Median:  3
* Mode:    3

...

Here's histogram of the distribution:

## Yellow Cards
Same for the yellow cards:

* Average per game: 3.2
* Median: 3
* Mode: 3

...

## Scatter Plot

So let's plot goals versus yellow cards. Each dot represents a game. We have to jitter the data so that we don't end up with overlapping dots. That way we get 380 dots. The drawback is that it looks like goals and yellow cards are continuous rather than discrete, which they obviously aren't.

As you can see there appears to be no correlation at all. Or actually if we look at the outliers, it seems the games with lots of yellow cards are not particularly high scoring, and the games with lots of goals seem to have low yellow cards. But remember the average is around 3 for both, so even in the outliers there's no strong trend.

## Linear Regression

Let's do a best fit line through the data and see if the slope is as flat as it looks like it should be after a quick eyeballing of the scatter. 

Yeah, pretty much flat.

## Last Ditch

OK, one last thing I want to have a look at to see if there is any trend whatsoever. I want to see what the average number of goals is for games where the yellow cards were above average and vice versa. Feel free to read that sentence again.

###### Average Goals in Games with 4 Yellow Cards or More

* 2.6 Compared to 2.7 for all games

###### Average Yellow Cards in Games with 3 Goals or More

* 3.2 Same as 3.2 for all games.

...

## Hypothesis Fail

So there you go, there appears to be absolutely no correlation between the number of goals in a game and the number of yellow cards. I'm sure we all feel much smarter after that exercise.  