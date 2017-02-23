Webscraping for Fun and Stats
Programmatically browsing for information
2016-11-20
## Webscraping

Webscraping is the art of programmatically browsing the web to retrieve information. It's fun. You get to feel a bit like a hacker, even though you probably don't actually want to hack into non-public facing websites to get data. And for a lot of publically available data, it is the most practical way to get it.

Webscraping is not web crawling. Web crawling is more open ended and basically just used to index websites, rather than extract information. Webscraping can well be looked down upon as an evil act, with connotations of DDoS and cracking. But it's also possible to scrape data from a website politely, without burdoning the webserver and breaching terms of use.

The hardest bit about web scraping is findning good information and tools. And that's where I hope this article can help. Many web scraping tools focus on interacting with the html returned from a HTTP request, but with the rise of modern javascript rendered websites, that approach is pretty limited, because increasingly the information you see on the page was sent with secondary requests, usually in the form of JSON or xml.

## The solution - Headless Web Browsing.

The solution to this is to use a headless web browser and browse the site programmatically. A headless web browser is just a browser without a graphical interface, so you pretty much have to use programming to control it. But it's easy. If you're reading this you've already done the hard part of finding decent info. If I do say so myself.

## The tools
We're going to use python for this. It's the goto language for such scripty things, but if your favourite programming language supports Selenium Web Driver, and it probably does, then feel free to use that instead. I hope the information in the rest of this post will still be useful to you. 

### Selenium Web Driver

First install Selenium web driver. Selenium web driver provides an interface for you to control a browser programmatically. Just open up your favourite search engine, probably Bing, and search for "install selenium webdriver \[your favourite programming language\] \[your operating system\]". Follow those instructions. This post isn't the place to list all of the combinations of architectures and programming languages.

### PhantomJS

Next you're going to have to install PhantomJS. PhantomJS is the actual headless web browser that you will be controlling with Selenium. You'll need NodeJS and the NPM to install it, so if you don't have Node installed, head over to [nodejs.org](https://nodejs.org) and follow the installaton instructions for your computer.

Once you've done that you can install PhantomJS from npm by typing the following into a terminal.

    npm install phantomjs -g 

## Acessing a page

Create a new python script **whatever_you_want_to_call_it.py** and open it up in your favourite text editor and enter the following:

    from selenium import webdriver
    import time

    driver = webdriver.PhantomJS()
    driver.set_window_size(1024, 768)
    driver.get('https://google.com')
    print 'Sleeping'
    time.sleep(3)
    driver.save_screenshot('screen.png')

Go ahead and run your script with

    python whatever_you_called_your_file.py

Then

    open screen.png

Wow, huh. pretty, pretty, pretty, pretty good. Consider that the "hello world" of webscraping. If you got it to work then you're ready to scrape the web for fun and stats. That little sleep in there, that's just to make sure all the javascript has rendered on the page. Not really a big deal when you're only hitting up google. 

## Inspecting with developer tools.
Right, I'm hoping you came here to find out how to get sports stats or some other publically available statistics, after reading my post [Does football follow the pareto principal?](/does-football-follow-pareto) So the first thing your going to want to do is head over to [BBC football results](http://www.bbc.com/sport/football/results) with chrome. I hope you're using chrome, if not, try to follow along anyway, but make sure you're using a browser with some sort of decent developer tools. For the record, I did not get my stats from the BBC, but I can't tell you where I got them. Even though sports statistics are public knowledge, scraping a website to steal, I mean scrape them, probably goes against some kind of terms of service and is most likely frowned upon.

Anyhoo, find a game that has a report link on the right side. Right click on the link and Inspect Element. I think that's what it's called, my chrome is in Danish. That will take you to the location of an anchor element for that link. Now in the element inspector, right click on the anchor element and select copy > copy selector.

Cool now we're ready to get the link for that programmatically.

## Retrieving information with selenium

So let's change our script to:

    from selenium import webdriver
    import time

    driver = webdriver.PhantomJS()
    driver.set_window_size(1024, 768)
    driver.get('https://google.com')
    print 'Sleeping'
    time.sleep(3)
    link = driver.find_element_by_css_selector('PASTE_YOUR_SELECTOR_HERE')
    print link.get_attribute('href')

Alternatively you could have copied the XPath from developer tools and used

    driver.find_element_by_xpath()

instead. I prefer css selectors just because I'm used to them from jquery.

So there you go, you just programmatically got some information from a webpage. Pretty sweet.
Now, you'll notice that the anchor tag there has a class of report, so what we could instead have done, is return a list of all the report anchors like so:

    links = driver.find_elements_by_css_selector('a.report')
    for link in links:
        print link.get_attribute('href')

Now we're talking. Notice that was find_element**s**_by_css_selector, that extra s in there gives us a list of all matches.

## Interacting with the page

So sometimes, we're going to have to interact with the page, such as clicking on things, or the bane of all webscrapers, scrolling down. Now a lot of tutorials online make scrolling down a web page sound way more complicated than it ought to be. It should in fact be this easy

    def scroll():
	   print 'Scrolling'
	   driver.execute_script("window.scrollTo(0, document.body.scrollHeight);")
	   time.sleep(5)

    for _ in range(12):
	    scroll()

That scrolls to the bottom of the page 12 times, waiting for 5 seconds each time to allow items to load. We've used driver.execute_script() to allow us to use javascript directly on the page. We can use that to click on elements as well. Say there was a tab called "stats" that didn't have a href and it was the third on a tab bar. We can get the selector from developer tools and click on the tab using the following script.

    statsbutton = driver.find_element_by_css_selector('div.tabs > ul > li:nth-child(3)')
    driver.execute_script("arguments[0].click();", statsbutton) 
    print 'Clicking'
    time.sleep(5)

again we add a sleep event just to allow any javascript to load. 5 seconds is probably a bit conservative, you can most likely go a bit shorter than that. 

## Looping through multiple web pages

So what if you scrape a page to get a whole bunch of urls and then you want to go to each of those pages and get some more information, for example go to each game of the season and collect stats. Well, my recommendation would be to first write your list of urls to file. Because scripting is sometimes a pain in the ass, and you might make a mistake. So you don't want to have to go to the front page every time you make a typo and trawl for links that you've already collected. So, to write to a file, simply:

    file = 'gamelinks.txt'
    target = open(file, 'w')

    for link in links: #links collected earlier in the script.
    target.write("%s\n" % link.get_attribute('href'))

Then you want to create a new script that goes through your text file and gets the stats. First get the links into a list.

    links = [link.rstrip('\n') for link in open('gamelinks.txt')]

Then you want to loop through the links. Here is the complete example script.

    from selenium import webdriver
    from selenium.common.exceptions import TimeoutException
    import time
    import io

    driver = webdriver.PhantomJS() # or add to your PATH
    driver.set_window_size(1024, 768) # optional

    links = [link.rstrip('\n') for link in open('fixtures.txt')]
    linksLength = len(links)

    for index, value in enumerate(links):
	    print 'Starting'
	    driver.set_page_load_timeout(60)
	    while True:
		    try:
			    driver.get(value)
		    except TimeoutException:
			    print "Timeout, retrying..."
			    continue
		    else:
			    break
	    print 'Sleeping'
	    time.sleep(3)
	    dataContainer = driver.find_element_by_css_selector('div.dataContainer')
	    dataJSON = dataContainer.get_attribute('data-game')
	    file = 'data{}.json'.format(index)
	    target = io.open(file, 'w', encoding="utf-8")
	    target.write(dataJSON)
	    target.close()
	    print 'File {} of {} written'.format(index, linksLength - 1)

So let's go through that. First the usual imports, plus selenium's TimeoutExeption. That's because if a page is too slow, it will crash your script and you'll have to do over, and we don't want that. We'll also import io for writing json, because footballers often have funny names that aren't ascii, so we need utf-8.

We make an enumerated for loop over the links so that we have access to the index for file naming and printing purposes. You'll notice that there's a lot of printing statements in there. I like to know what's going on while my scripts are running. Web scraping scripts can be excruciatingly slow, so those print statements are a sanity check.

We do a try except on the driver.get so that we can handle pages that are taking too long.

So in this example we just happened to stumble upon a website that stores it's data in json format right in a data attribute on the page. Yes this awesomeness does sometimes happen in the real world. In fact, that's how I got the premier league stats from an unidentified website. The rest of the script saves each game's json object to it's own file. Again this is for ease of working with. Also, there are 380 games in a premier league season, if you're script craps out after 250 games or something, you really don't want to start from that start again, as this could take hours.

## Awesome

So there you go, that's how to scrape a whole bunch of javascipt rendered webpages that follow a similar template. Easy peasy. It's slow, much slower that pinging the server for the html page, but that's not going to get you much in the days of the single page app.

## What now?

So now that we have a whole bunch of JSON files, what are we going to do with them? Well, I suggest storing your data in an SQLite database, and that's what my next post is all about. So join me to see [how I made an SQLite file](/json-sqlite-csv) that contains all of the stats from the 2015/16 English Premier League season. 