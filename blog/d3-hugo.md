Embedding a D3 graph in a Hugo website
It's super easy
2016-11-25
## Hugo and D3

I've seen a couple of posts on forums asking how to get a D3 graph or animation using a static site generator like Hugo. It's actually very simple. You see, you can put HTML right in along side your markdown.

## Adding a Container div

So what you want to do is, in your markdown file, create a Div with an ID where you want to put the graph. Like this:

    ## Some smart heading
    blah blah blah, interesting fact nugget about data
    and here's a graph that shows off my wisdom
    <div id="myGraph"></div>

## Attaching to your div in your javascript

So then, in your javascript you want to do the following:

    var svg = d3.select("#myGraph")
	.append("svg")
	....

## Saving your javascript in a separate file

One caveat. You'll have to save your javascript as a separate file, because otherwise the markdown processor will escape out all the important javascript stuff and your graph won't work. Most tutorials for D3 you see online have CSS and Javascript inline with the HTML, nuh uh. Don't do that.

## Loading D3 first.
So at the end of your markdown file, you'll want to link to two scripts like this:

    <script src="https://d3js.org/d3.v4.min.js"></script>
	<script src="/mygraph.js"></script>

You'll need to load the D3 library first for your graph to render. No webpack here. Make sure to save your javascript file in your static folder. If you have a lot of static files, you may want to keep it in a separate js folder and then your script link will look like this:

    <script src="/js/mygraph.js"></script>

So there you go, it's very easy, the main things to consider are, keeping your js and css separate from your html, and creating the container div and script links as html directly in your markdown. I hope this helps.