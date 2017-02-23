Does Football Follow the Pareto Principal
A look at top goal scorers versus the rest of the pack
2016-11-21
Let's keep this first post nice and basic.

*If you're intested in data science, you may be interested in the accompanying posts [Web scraping for fun and stats](/webscraping-fun-stats), [JSON to SQLite to csv - Data transforations](/json-sqlite-csv) and [D3: A CDF graph with hover tooltips](/D3-CDF-with-hover)*
## The Pareto principal
Vilfredo Pareto was an Italian polymath who introduced some important concepts in microeconomics. But he's mostly famous for his most meme-able observation: the 80/20 rule. He noticed that in his time, about 20% of Italians owned about 80% of the land. Since then people have noticed this 80/20 rule all over the place. 20% of your customers make up 80% of your sales. 20% of the global population have 80% of the wealth, 20% of your friends are having 80% of the sex. etc. etc. 

## Football

But our question is, does football fit this pattern? Specifically, I hypothesise that 20% of the players score 80% of the goals. And by players I mean anyone who at least made it to the bench, so no reserves. For our sample I'll be using the last complete season of the English Premier League - 2015-16. 

So well, here's our graph.

<div id="paretoGraph"></div>

You can hover over the individual data points of the bottom line to see the players and how many goals they scored. It's a little cramped with so many data points.

The bottom line represents the percentage of goals each player has scored whereas the top line represents the cumulation of the goals. The bottom line can easily mislead you into thinking that scoring is reasonably equally distributed, but the cumulation line you can clearly see that infact:

## It's True!
Around 20% of the players do indeed score around 80% of the goals.

## What's next
So that was a really basic alternative way to view some common football stats. I plan to follow up with more football related articles soon, especially seeing as I have a very detailed football data set. More about that [here](/webscraping-fun-stats).

<script src="https://d3js.org/d3.v4.min.js"></script>
<script src="/paretoscorers.js"></script>