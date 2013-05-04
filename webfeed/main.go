// +build ignore

package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/velour/feedme/webfeed"
)

func main() {
	fmt.Println("Atom")
	feed, err := webfeed.Read(strings.NewReader(atomData))
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(feed)

	fmt.Println("Rss")
	feed, err = webfeed.Read(strings.NewReader(rssData))
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(feed)
}

var atomData = `
<?xml version="1.0" encoding="utf-8" standalone="yes"?>
<feed xmlns="http://www.w3.org/2005/Atom">
<title>Programming in the 21st Century</title>
<link rel="self" href="http://prog21.dadgum.com/atom.xml"/>
<link rel="alternate" href="http://prog21.dadgum.com/"/>
<id>http://prog21.dadgum.com/</id>
<updated>2013-04-08T00:00:00-06:00</updated>
<entry>
<title>Exploring the Lower Depths of Terseness</title>
<link rel="alternate" type="text/html" href="http://prog21.dadgum.com/170.html"/>
<id>http://prog21.dadgum.com/170.html</id>
<published>2013-04-08T00:00:00-06:00</published>
<updated>2013-04-08T00:00:00-06:00</updated>
<author><name>James Hague</name></author>
<content type="xhtml"><div xmlns="http://www.w3.org/1999/xhtml">There's a 100+ year old system for recording everything that happens in a baseball game. It uses sheet of paper with a small box for each batter. Whether that batter gets to base or is out--and why--gets coded into that box. It's a scorekeeping method that's still in use at the professional and amateur level, and at major league games you can buy a program which includes a scorecard.
<br/><br/>What's surprising is how cryptic the commonly used system is. For starters, each position is identified by a number. The pitcher is 1. The center fielder 8. If the ball is hit to the shortstop who throws it to the first baseman, the sequence is 6-3. See, there isn't even the obvious mnemonic of the first, second, and third basemen being numbered 1 through 3 (they're 3, 4, and 5).
<br/><br/>In programming, no one would stand for this. It breaks the rule of not having magic numbers. I expect the center fielder would be represented by something like:
<pre>visitingTeam.outfield.center
</pre>The difference, though, is that programming isn't done in real-time like scorekeeping. After the initial learning curve, 8 is much more concise, and the terseness is a virtue when recording plays with the ball moving between multiple players. Are we too quick to dismiss extremely terse syntax and justify long-winded notations because they're easier for the uninitiated to read?
<br/><br/>Suppose you have a file where each line starts with a number in parentheses, like "(124)", and you want to replace that number with an asterisk. In the vim editor the keystrokes for this are "<code>^cib*</code>" followed by the escape key. "^" moves to the start of the line. The "c" means you're going to change something, but what? The following "ib" means "inner block" or roughly "whatever is inside parentheses." The asterisk fills in the new character.
<br/><br/>Once you get over the dense notation, you may notice a significant win: this manipulation of text in vim can be described and shared with others using only five characters. There's no "now press control+home" narrative.
<br/><br/>The ultimate in terse programming languages is J. The boring old "*" symbol not only multiplies two numbers, but it pairwise multiplies two lists together (as if a map operation were built in) and also multiplies a scalar value with each element of a list, depending on the types of its operands. 
<br/><br/>That's what happens with two operands anyway. Each verb (the J terminology for "operator"), also works in a unary fashion, much like the minus sign in C represents both subtraction and negation. When applied to a lone value "*" is the sign function, returning either -1, 0, or 1 if the operand is negative, zero, or positive.
<br/><br/>So now each single-character verb has two meanings, but it goes further than that. To increase the number of symbolic verbs, each can have either a period or a colon as a second character, and then each of these have both one and two operand versions. "*:" squares a single parameter or returns the nand ("not and") of two parameters. Then there's the two operand version of "*." which computes the least common multiple, and I'll give it up now before everyone stops reading.
<br/><br/>Here's the reason for this madness: it allows a wide range of built-in verbs that never conflict with user-defined, alphanumeric identifiers. Without referencing a single library you've got access to prime number generation ("p:"), factorial ("!"), random numbers ("?"), and matrix inverse ("%.").
<br/><br/>Am I recommending that you switch to vim for text editing and J for coding? No. But when you see an expert working with those tools, getting results with fewer keystrokes than it would take to import a Python module, let alone the equivalent scripting, then...well, there's something to the terseness that's worth remembering. It's too impressive to ignore simply because it doesn't line up with the prevailing aesthetic for readable code.
<br/><br/>(If you liked this, you might enjoy <a href="http://prog21.dadgum.com/114.html">Papers from the Lost Culture of Array Languages</a>.)
</div></content>
</entry>
<entry>
<title>Expertise, the Death of Fun, and What to Do About It</title>
<link rel="alternate" type="text/html" href="http://prog21.dadgum.com/169.html"/>
<id>http://prog21.dadgum.com/169.html</id>
<published>2013-03-12T00:00:00-06:00</published>
<updated>2013-03-12T00:00:00-06:00</updated>
<author><name>James Hague</name></author>
<content type="xhtml"><div xmlns="http://www.w3.org/1999/xhtml">I've started writing this twice before. The first time it turned into <a href="http://prog21.dadgum.com/167.html">Simplicity is Wonderful, But Not a Requirement</a>. The second time it ended up as <a href="http://prog21.dadgum.com/168.html">Don't Be Distracted by Superior Technology</a>. If you re-read those you might see bits and pieces of what I've been wanting to say, which goes like this:
<br/><br/>There is danger in becoming an expert. Long-term exposure to programming, coding, software development--whatever you want to call it--changes you. You start to recognize the extreme complexity in situations where there doesn't need to be any, and it eats at you. You realize how broken the tools are. You discover bygone flashes of amazing beauty in old systems that have been set aside in favor of the way things have always been done.
<br/><br/>This is a bad line of thinking.
<br/><br/>It's why you run into twenty-year veteran coders who can no longer write <a href="http://imranontech.com/2007/01/24/using-fizzbuzz-to-find-developers-who-grok-coding/">FizzBuzz</a>. It's why people right out of school often create impressive and impossible-seeming things, because they haven't yet developed an aesthetic that labels all of that successful hackery as ugly and wrong. It's why some programmers migrate to more and more obscure languages, trading productivity for poetic tinkering.
<br/><br/>Maybe a better title for this piece is "So You've Become Jaded and Dissatisfied. Now What?"
<br/><br/><b>Cultivate a "try it first" attitude.</b> Yes, it's amusing to read about those silly developers who can't write FizzBuzz. But your first reaction should be to set aside the article and <i>try implementing it yourself</i>. 
<br/><br/><b>Active learning or bust.</b> Don't bother with tutorials or how-to books unless you're going to use the information immediately. Fire up your favorite interpreter and play along as you read. Don't take the author's word for anything; prove it to yourself. Do the exercises and invent your own.
<br/><br/><b>Be realistic about the limitations of your favorite programming language.</b> I enjoy Erlang, but it's <a href="http://prog21.dadgum.com/38.html">puzzle language</a>, meaning that some truly trivial problems don't have a straightforward mapping to the strengths of the language (such as most algorithms based around destructive array updates). When I don't have a clear picture of what features I'm going to need, I reach for something with few across-the-board sticking points, like Python. Sometimes the cleanest approach involves straightforward loops and counters and <code>return</code> statements right there in the middle of it all.
<br/><br/><b>Let ease of implementation trump perfection.</b> Yes, yes, grepping an XML file is fundamentally wrong...but it often works and is easier than dealing with an XML parsing library. Yes, you're supposed to properly handle exceptions that get thrown and always check for memory errors and division by zero and all that. If you're writing code for a manned Mars mission, then <i>please</i> take the time to do it right. But for a personal tool or a prototype, it's okay if you don't. Really. It's better to focus on the fun parts and generating useful results quickly.
</div></content>
</entry>
<entry>
<title>Don't Be Distracted by Superior Technology</title>
<link rel="alternate" type="text/html" href="http://prog21.dadgum.com/168.html"/>
<id>http://prog21.dadgum.com/168.html</id>
<published>2013-03-03T00:00:00-06:00</published>
<updated>2013-03-03T00:00:00-06:00</updated>
<author><name>James Hague</name></author>
<content type="xhtml"><div xmlns="http://www.w3.org/1999/xhtml">Not long after I first learned C, I stumbled across a lesser-used language called Modula-2. It was designed by Niklaus Wirth who previously created Pascal. While Pascal was routinely put down for being awkwardly restrictive, Wirth nudged and reshaped the language into Modula-2, arguably the finest systems-level programming language of its day.
<br/><br/>Consider that Modula-2 had the equivalent of C++ references from the start (and for the record, so did Pascal and ALGOL). Most notably, if you couldn't guess from the name, Modula-2 had a true module system that has managed to elude the C and C++ standards committees for decades.
<br/><br/>My problem became that I had been exposed to the Right Approach to separately compiled modules, and going back to C felt backward--even broken.
<br/><br/>When I started exploring functional programming, I used to follow the Usenet group comp.lang.functional. A common occurrence was that someone struggling with how to write programs free of destructive updates would ask a question. As the thread devolved into bickering about type systems and whatnot, someone would inevitably point out a language that gracefully handled that particular issue in a much nicer way than Haskell, ML, or Erlang.
<br/><br/>Except that the suggested language was an in-progress research project.
<br/><br/>The technology world is filled with cases where smart and superior alternatives exist, but their existence makes no difference because you can't use them. 1980s UNIX was incredibly stable compared to MS-DOS, but it was irrelevant if you intended to  use MS-DOS software. Clojure and Factor are wonderful languages, but if you want to write iOS games then you're better off pretending you've never heard of them. Not only are they not good options for iOS, at least not at the moment, but going so against the grain brings extra work and headaches with it.
<br/><br/>Words like "better," "superior," and "right," are misleading. Yes, Modula-2 has a beautiful module system, but that's negated by being a fringe language that isn't likely to be available from the start when exciting new hardware is released. Erlang isn't as theoretically beautiful as those cutting-edge research languages, but it's been through the forge of shipping large-scale systems. What may look like warts upon first glance may be the result of pragmatic choices.
<br/><br/>There's much more fun to be had building things than constantly being distracted by better technology.
<br/><br/>(If you liked this, you might enjoy <a href="http://prog21.dadgum.com/140.html">The Pace of Technology is Slower than You Think</a>.)
</div></content>
</entry>
</feed>
`

var rssData = `
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" media="screen" href="/~d/styles/rss2full.xsl"?><?xml-stylesheet type="text/css" media="screen" href="http://feeds.arstechnica.com/~d/styles/itemcontent.css"?><rss xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:wfw="http://wellformedweb.org/CommentAPI/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/" xmlns:slash="http://purl.org/rss/1.0/modules/slash/" xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" version="2.0">

<channel>
	<title>Ars Technica » Risk Assessment</title>
	
	<link>http://arstechnica.com</link>
	<description>The Art of Technology</description>
	<lastBuildDate>Sat, 27 Apr 2013 19:58:08 +0000</lastBuildDate>
	<language>en-US</language>
	<sy:updatePeriod>hourly</sy:updatePeriod>
	<sy:updateFrequency>1</sy:updateFrequency>
	<generator>http://wordpress.org/?v=3.5.1</generator>
		<atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="self" type="application/rss+xml" href="http://feeds.arstechnica.com/arstechnica/security" /><feedburner:info uri="arstechnica/security" /><atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="hub" href="http://pubsubhubbub.appspot.com/" /><item>
		<title>Why LivingSocial’s 50-million password breach is graver than you may think</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/dwzVGAxVAf0/</link>
		<comments>http://arstechnica.com/security/2013/04/why-livingsocials-50-million-password-breach-is-graver-than-you-may-think/#comments</comments>
		<pubDate>Sat, 27 Apr 2013 19:00:49 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[Crackers]]></category>
		<category><![CDATA[Cracking]]></category>
		<category><![CDATA[passwords]]></category>
		<category><![CDATA[security]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=255413</guid>
		<description><![CDATA[No, cryptographically scrambled passwords are <em>not</em> hard to decode.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/warning.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.geograph.ie/profile/12869">Rossographer</a>
				</div>
	  </div>
  </div>
 <p>LivingSocial.com, a site that offers daily coupons on restaurants, spas, and other services, has suffered a security breach that has exposed names, e-mail addresses and password data for up to 50 million of its users. If you're one of them, you should make sure this breach doesn't affect other accounts that may be impacted.</p>
<p>In an e-mail sent Friday, CEO Tim O'Shaughnessy told customers the stolen passwords had been hashed and salted. That means passcodes were converted into one-way cryptographic representations that used random strings to cause each hash string to be unique, even if it corresponded to passwords chosen by other LivingSocial users. He went on to say "your Living Social password would be difficult to decode." This is a matter for vigorous debate, and it very possibly could give users a false sense of security.</p>
<p>As Ars <a href="http://arstechnica.com/security/2012/08/passwords-under-assault/">explained before</a>, advances in hardware and hacking techniques make it trivial to crack passwords that are presumed strong. LivingSocial engineers should be applauded for adding cryptographic salt, because the measure requires password cracking programs to guess the plaintext for each individual hash, rather than guessing passwords for millions of tens of millions of hashes all at once. But a far more important measure of protection, password cracking experts say, is the hashing algorithm used. SHA1, the <a href="https://www.livingsocial.com/createpassword">algorithm used by LivingSocial</a>, is an extremely poor choice for secure password storage. Like MD5 and even the newly adopted SHA3 algorithms, it's designed to operate quickly and with a minimal amount of computing resources. A far better choice would have been bcrypt, scrypt, or PBKDF2.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/why-livingsocials-50-million-password-breach-is-graver-than-you-may-think/#p3">Read 4 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/why-livingsocials-50-million-password-breach-is-graver-than-you-may-think/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/dwzVGAxVAf0" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/why-livingsocials-50-million-password-breach-is-graver-than-you-may-think/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/why-livingsocials-50-million-password-breach-is-graver-than-you-may-think/</feedburner:origLink></item>
		<item>
		<title>Police arrest suspect accused of “unprecedented” DDoS attack on Spamhaus</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/63zTw8vTR3g/</link>
		<comments>http://arstechnica.com/security/2013/04/police-arrest-spamhaus-attacker-accused-of-unprecedented-ddos-attack/#comments</comments>
		<pubDate>Fri, 26 Apr 2013 18:45:51 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[cyberbunker]]></category>
		<category><![CDATA[DDoS]]></category>
		<category><![CDATA[distributed denial of service]]></category>
		<category><![CDATA[spamhaus]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=254757</guid>
		<description><![CDATA[Is suspect identified as "SK" actually CyberBunker affiliate Sven Olaf Kamphuis? Likely.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>Spanish authorities have arrested a 35-year-old Dutchman they say is "suspected of unprecedented heavy attacks" on Spamhaus, the international group that helps network owners around the world block spam.</p>
<p>A <a href="http://www.om.nl/actueel/nieuws-persberichten/@160856/nederlander/">press release</a> (<a href="http://www.microsofttranslator.com/bv.aspx?ref=SERP&amp;br=ro&amp;mkt=en-US&amp;dl=en&amp;lp=NL_EN&amp;a=http%3a%2f%2fwww.om.nl%2factueel%2fnieuws-persberichten%2f%40160856%2fnederlander%2f">English translation here</a>) issued by the Dutch Public Prosecutor Service identified the suspect only by the initials SK and said he was living in Barcelona. A variety of circumstantial evidence, mostly taken from <a href="https://www.facebook.com/cb3rob">this Facebook profile</a>, strongly suggests the suspect is one Sven Olaf Kamphuis. He's the man quoted in a <a href="http://www.nytimes.com/2013/03/27/technology/internet/online-dispute-becomes-internet-snarling-attack.html?pagewanted=all&amp;_r=0">March 26 <em>New York Times</em> article</a> saying a Dutch hosting company called CyberBunker, which Kamphuis is affiliated with, was behind distributed denial-of-service attacks aimed at Spamhaus. Kamphuis <a href="https://www.facebook.com/cb3rob/posts/348246075277329">later denied</a> he or CyberBunker had anything to do with the attacks.</p>
<p>With peaks of 300 gigabits per second, the March attacks were among the biggest ever recorded. Besides their size, they were also notable because they attacked the London Internet Exchange, a regional hub where multiple networks from different service providers connect. As <a href="http://arstechnica.com/author/peter-bright/">Ars writer Peter Bright</a> explained, the size and technique threatened to <a href="http://arstechnica.com/security/2013/03/spamhaus-ddos-grows-to-internet-threatening-size/">clog up the Internet's core infrastructure</a> and make access to the rest of the Internet slow or impossible. While some critics said that assessment was overblown, Bright provided <a href="http://arstechnica.com/security/2013/04/can-a-ddos-break-the-internet-sure-just-not-all-of-it/">this follow-up</a> explaining why the attacks had the potential to break key parts of the Internet.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/police-arrest-spamhaus-attacker-accused-of-unprecedented-ddos-attack/#p3">Read 2 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/police-arrest-spamhaus-attacker-accused-of-unprecedented-ddos-attack/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/63zTw8vTR3g" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/police-arrest-spamhaus-attacker-accused-of-unprecedented-ddos-attack/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/police-arrest-spamhaus-attacker-accused-of-unprecedented-ddos-attack/</feedburner:origLink></item>
		<item>
		<title>Google bans self-updating Android apps, possibly including Facebook’s</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/ClHRmc0eWN4/</link>
		<comments>http://arstechnica.com/information-technology/2013/04/google-bans-self-updating-android-apps-possibly-including-facebooks/#comments</comments>
		<pubDate>Fri, 26 Apr 2013 15:05:55 +0000</pubDate>
		<dc:creator>Jon Brodkin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[android]]></category>
		<category><![CDATA[Facebook]]></category>
		<category><![CDATA[google play]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=254545</guid>
		<description><![CDATA[Facebook end run around Google Play store followed by change in Google policy.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>About six weeks ago, users of Facebook's Android application <a href="http://liliputing.com/2013/03/facebook-pushes-android-update-to-enable-silent-updates-bypassing-the-play-store.html">noticed</a> that they were being asked to install a new version—without going to the Google Play app store.</p>
<p>Android is far more permissive than iOS regarding the installation of third-party applications, even allowing installation from third-party sources if the user explicitly allows it. However, it's unusual for applications delivered through the official Google store to receive updates outside of the store's updating mechanism.</p>
<p>Google has now changed the Google Play store polices in an apparent attempt to avoid Facebook-like end runs around store-delivered updates. Under the "Dangerous Products" section of the <a href="https://play.google.com/about/developer-content-policy.html">Google Play developer policies</a>, Google now states that "[a]n app downloaded from Google Play may not modify, replace or update its own APK binary code using any method other than Google Play's update mechanism." A <a href="http://www.droid-life.com/2013/04/25/google-updates-play-store-content-policy-to-remind-developers-they-cant-update-apks-except-with-googles-update-mechanism-stares-directly-at-facebook/">Droid-Life</a> article says the language update occurred Thursday. APK (standing for application package file) is the file format used to install applications on Android.</p>
</div><p><a href="http://arstechnica.com/information-technology/2013/04/google-bans-self-updating-android-apps-possibly-including-facebooks/#p3">Read 4 remaining paragraphs</a> | <a href="http://arstechnica.com/information-technology/2013/04/google-bans-self-updating-android-apps-possibly-including-facebooks/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/ClHRmc0eWN4" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/information-technology/2013/04/google-bans-self-updating-android-apps-possibly-including-facebooks/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/information-technology/2013/04/google-bans-self-updating-android-apps-possibly-including-facebooks/</feedburner:origLink></item>
		<item>
		<title>Potent DDoS attacks on Mt. Gox delay rollout of new virtual currency</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/cMrkW42Up8o/</link>
		<comments>http://arstechnica.com/security/2013/04/potent-ddos-attacks-on-mt-gox-delays-rollout-of-new-virtual-currency/#comments</comments>
		<pubDate>Thu, 25 Apr 2013 16:48:42 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=253943</guid>
		<description><![CDATA[Support of Litecoin is postponed as Bitcoin exchange struggles to stay online.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/explosion-640x365.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="https://commons.wikimedia.org/wiki/File:USAF_EOD_explosion.jpg">Wikimedia Commons</a>
				</div>
	  </div>
  </div>
 <p>Mt. Gox, the world's largest Bitcoin exchange, is delaying plans to support a new form of virtual currency known as <a href="http://litecoin.org/">Litecoin</a> following a series of debilitating Internet attacks that are growing increasingly powerful.</p>
<p>The most recent distributed denial-of-service (DDoS) attack to hit Mt. Gox came on Sunday, and it knocked the Tokyo-based exchange offline for four hours, officials said in a <a href="https://mtgox.com/pdf/20130424_ddos_statement_and_faq.pdf">statement issued Wednesday</a>. Unlike more traditional DDoS attacks, which flood websites' routers and servers with more junk data than they can handle, the latest assault targeted Web applications the Mt. Gox site uses to process and secure customer transactions. That's known as <a href="https://en.wikipedia.org/wiki/Application_Layer">Layer 7, or the application layer</a>, of the networking stack.</p>
<p>"What we are experiencing lately are 'Layer 7' DDoS attacks," the statement read. "Unlike your average DDoS (which overloads the servers with traffic to the sites as a whole) these are much more creative and harder to detect in that they target specific elements of the site and make it difficult to distinguish malicious traffic from normal traffic. The attackers' goal is to shut down the exchange, either thorough the DDoS itself, or by forcing Mt. Gox to take measures that have the same effect."</p>
</div><p><a href="http://arstechnica.com/security/2013/04/potent-ddos-attacks-on-mt-gox-delays-rollout-of-new-virtual-currency/#p3">Read 6 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/potent-ddos-attacks-on-mt-gox-delays-rollout-of-new-virtual-currency/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/cMrkW42Up8o" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/potent-ddos-attacks-on-mt-gox-delays-rollout-of-new-virtual-currency/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/potent-ddos-attacks-on-mt-gox-delays-rollout-of-new-virtual-currency/</feedburner:origLink></item>
		<item>
		<title>Critical app flaw bypasses screen lock on up to 100 million Android phones</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/ssiluDZW-kA/</link>
		<comments>http://arstechnica.com/security/2013/04/crital-app-flaw-bypasses-screen-lock-on-up-to-100-million-android-phones/#comments</comments>
		<pubDate>Wed, 24 Apr 2013 19:15:19 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Gear & Gadgets]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[smartphone security]]></category>
		<category><![CDATA[viber android]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=253291</guid>
		<description><![CDATA[Skype-rival Viber confirms the bug allows hackers to take control of locked devices.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/viber-flaw-demo-640x377.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="https://www.youtube.com/watch?feature=player_embedded&amp;v=gSB1MvGFfB4">Bkav Internet Security</a>
				</div>
	  </div>
  </div>
 <p>A critical flaw in an Android app downloaded as many as 100 million times allows attackers to take full control of handsets even when they're protected by screen locks.</p>
<p>The vulnerability in the <a href="https://play.google.com/store/apps/details?id=com.viber.voip&amp;feature=search_result#?t=W251bGwsMSwxLDEsImNvbS52aWJlci52b2lwIl0.">Skype rival known as Viber</a> affects Android smartphone brands such as Samsung, Sony, and HTC, according to a <a href="http://www.bkav.com/top-news/-/view_content/content/46264/critical-flaw-in-viber-allows-full-access-to-android-smartphones-bypassing-lock-screen">blog post</a> published Tuesday by Bkav Internet Security. Although attack techniques differ from model to model, they all exploit programming logic in the way Viber handles popup messages, researchers with the company wrote.</p>
<p>A spokesman Viber Media, maker of the affected app, said company officials learned of the vulnerability on Wednesday and plan to release a fix next week.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/crital-app-flaw-bypasses-screen-lock-on-up-to-100-million-android-phones/#p3">Read 3 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/crital-app-flaw-bypasses-screen-lock-on-up-to-100-million-android-phones/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/ssiluDZW-kA" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/crital-app-flaw-bypasses-screen-lock-on-up-to-100-million-android-phones/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/crital-app-flaw-bypasses-screen-lock-on-up-to-100-million-android-phones/</feedburner:origLink></item>
		<item>
		<title>Hacked AP Twitter feed reporting fake White House attack rocks markets</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/fkdLxkc9Y1Q/</link>
		<comments>http://arstechnica.com/security/2013/04/hacked-ap-twitter-feed-rocks-market-after-sending-false-news-flash/#comments</comments>
		<pubDate>Tue, 23 Apr 2013 19:44:03 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[hacking]]></category>
		<category><![CDATA[Twitter]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=251967</guid>
		<description><![CDATA[Account compromise comes after AP targeted by malware and phishing e-mails.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/dow-drop-640x368.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">The seven-minute drop in the Dow Jones Industrial Average touched off by a single tweet falsely claiming the White House had been bombed. It temporarily wiped out about 1 percent of the average, which can translate into millions or billions of dollars in market capitalization.</div>
	
	  </div>
  </div>
 <p>Stock prices plunged and then quickly recovered after a Twitter account belonging to the Associated Press was hacked and used to send a bogus report falsely claiming that the White House had been bombed and President Obama was injured.</p>
<p>"The @AP Twitter account has been suspended after it was hacked," an unaffected Twitter account belonging to the news organization <a href="https://twitter.com/APStylebook/status/326749129453752320">confirmed</a>. "The tweet about an attack on the White House was false."</p>
<p>In a testament to the power that social media has on real-world finances, the Dow Jones Industrial Average fell 150 points, or about 1 percent, immediately following the tweet, with other indexes reacting similarly. The Dow quickly regained the lost ground about seven minutes after the sell-off began, when the AP confirmed that the report was false.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/hacked-ap-twitter-feed-rocks-market-after-sending-false-news-flash/#p3">Read 5 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/hacked-ap-twitter-feed-rocks-market-after-sending-false-news-flash/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/fkdLxkc9Y1Q" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/hacked-ap-twitter-feed-rocks-market-after-sending-false-news-flash/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/hacked-ap-twitter-feed-rocks-market-after-sending-false-news-flash/</feedburner:origLink></item>
		<item>
		<title>Java users beware: Exploit circulating for just-patched critical flaw</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/XPwZ1_TjTMs/</link>
		<comments>http://arstechnica.com/security/2013/04/java-users-beware-exploit-circulating-for-just-patched-critical-flaw/#comments</comments>
		<pubDate>Tue, 23 Apr 2013 16:27:40 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[computer security]]></category>
		<category><![CDATA[exploit]]></category>
		<category><![CDATA[Java]]></category>
		<category><![CDATA[malware]]></category>
		<category><![CDATA[oracle]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=251799</guid>
		<description><![CDATA[If you haven't installed last week's Java update, now would be a good time.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/java1.jpg">
</div>
 <p>If you haven't installed <a href="http://www.oracle.com/technetwork/topics/security/javacpuapr2013-1928497.html">last week's patch from Oracle</a> that plugs dozens of critical holes in its Java software framework, now would be a good time. As in immediately. As in, really, <em>right now</em>.</p>
<p>In the past few days, attack code targeting one of the many remote-code-execution vulnerabilities fixed in Java 7 Update 21 was folded into either the <a href="https://twitter.com/TimoHirvonen/status/326682538942803970">folded into the RedKit</a> or <a href="https://twitter.com/TimoHirvonen/status/326708970117013505">CrimeBoss</a> exploit kit. By Sunday, that attack code was being actively unleashed on unsuspecting end users, according to a <a href="http://www.f-secure.com/weblog/archives/00002544.html">short blog post</a> published by a researcher from antivirus provider F-Secure.</p>
<p>The post doesn't say where the attacks were being hosted or precisely how attackers are using them. Still, Oracle describes <a href="http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2013-2423">the vulnerability</a> as allowing remote code execution without authentication. And that means you should install the patch before you do anything else today. The track record of malware purveyors of abusing advertising networks, <a href="http://arstechnica.com/security/2013/04/exclusive-ongoing-malware-attack-targeting-apache-hijacks-20000-sites/">compromised Apache servers</a>, and other legitimate enterprises means readers could encounter attacks even when they're browsing a site they know and trust.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/java-users-beware-exploit-circulating-for-just-patched-critical-flaw/#p3">Read 3 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/java-users-beware-exploit-circulating-for-just-patched-critical-flaw/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/XPwZ1_TjTMs" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/java-users-beware-exploit-circulating-for-just-patched-critical-flaw/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/java-users-beware-exploit-circulating-for-just-patched-critical-flaw/</feedburner:origLink></item>
		<item>
		<title>More “BadNews” for Android: New malicious apps found in Google Play</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/xfIujOh_WVo/</link>
		<comments>http://arstechnica.com/security/2013/04/more-badnews-for-android-new-malicious-apps-found-in-google-play/#comments</comments>
		<pubDate>Mon, 22 Apr 2013 19:05:36 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Gear & Gadgets]]></category>
		<category><![CDATA[Risk Assessment]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=250987</guid>
		<description><![CDATA[The code family used to push malware circulated as early as June 2012.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/android.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="https://secure.flickr.com/photos/quinnanya/4754719107/sizes/z/in/photostream/">quinn.anya</a>
				</div>
	  </div>
  </div>
 <p>The family of Android malware that slipped past security defenses and infiltrated <a href="https://play.google.com/">Google Play</a> is more widespread than previously thought. New evidence shows it was folded into three additional apps and has been operating for at least 10 months, according to security researchers.</p>
<p>BadNews, as the malicious ad network library is called, has been included in at least 35 different apps that were available on Google servers for download, researchers from antivirus provider Bitdefender <a href="http://www.hotforsecurity.com/blog/badnews-android-malware-active-since-june-2012-5998.html">said Monday</a>. As <a href="http://arstechnica.com/security/2013/04/family-of-badnews-malware-in-google-play-downloaded-up-to-9-million-times/">Ars reported last week</a>, figures provided by Google showed they had been downloaded anywhere from two million to nine million times. Although Google had removed 32 apps as of Friday, company security personnel didn't remove the additional three apps until they were flagged this weekend by Bitdefender. Apps that contain the BadNews code upload phone numbers, unique device identifiers, and other data from infected phones and then present end users with prompts to download and install fake updates for legitimate applications such as Skype.</p>
<p>The Bitdefender report came as researchers from security firm Fortinet reported the deactivation of a <a href="http://blog.fortinet.com/Secure-malware--/">Google Play developer account that was also pushing a suspicious app</a>.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/more-badnews-for-android-new-malicious-apps-found-in-google-play/#p3">Read 7 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/more-badnews-for-android-new-malicious-apps-found-in-google-play/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/xfIujOh_WVo" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/more-badnews-for-android-new-malicious-apps-found-in-google-play/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/more-badnews-for-android-new-malicious-apps-found-in-google-play/</feedburner:origLink></item>
		<item>
		<title>Family of “BadNews” malware in Google Play downloaded up to 9 million times</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/hS0_oWvBHPU/</link>
		<comments>http://arstechnica.com/security/2013/04/family-of-badnews-malware-in-google-play-downloaded-up-to-9-million-times/#comments</comments>
		<pubDate>Sat, 20 Apr 2013 15:21:42 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Gear & Gadgets]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[GooglePlay]]></category>
		<category><![CDATA[malware]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=249891</guid>
		<description><![CDATA[Apps steal sensitive data, push SMS app that racks up charges to pricey service.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/zombie-android.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="https://secure.flickr.com/photos/21986855@N07/">greyweed</a>
				</div>
	  </div>
  </div>
 <p>Security researchers have unearthed a family of malware for Android-based smartphones that has been downloaded as many as 9 million times from <a href="https://play.google.com/">Google Play</a>, the official distribution platform hosted on Google servers.</p>
<p>BadNews, as the library of malicious code has been dubbed, was folded in to at least 32 applications offered by four different developer accounts, according to a <a href="https://blog.lookout.com/blog/2013/04/19/the-bearer-of-badnews-malware-google-play/">blog post</a> published Friday by Android app provider Lookout Mobile Security. Handsets that run the poisoned apps connect to a rogue server every four hours and report several pieces of sensitive information, including the device phone number and its unique serial number, known as an <a href="https://en.wikipedia.org/wiki/IMEI">International Mobile Station Equipment Identity</a>. The command and control servers, which were still operational as of Friday, also force some phones to display prompts to install <a href="https://blog.lookout.com/blog/2012/04/11/the-continuing-saga-of-fake-app-toll-fraud/">AlphaSMS</a>, a trojan that racks up charges by sending text messages to pricey services.</p>
<p>The people behind the campaign were able to sneak BadNews past Google defenses by adding the malware library to innocuous apps after they had already been submitted to Google Play. That gave the appearance of trustworthiness to measures such Bouncer, the <a href="http://arstechnica.com/business/2012/02/at-long-last-malware-scanning-comes-to-googles-android-market/">cloud-based service that scours Play for abusive apps</a>. It was only later that the apps were updated to carry out the attacks. Figures provided by Google Play showed the targeted apps had been downloaded from 2 million to 9 million times. It's unclear how many of the downloads involved apps after they had been updated to include BadNews.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/family-of-badnews-malware-in-google-play-downloaded-up-to-9-million-times/#p3">Read 5 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/family-of-badnews-malware-in-google-play-downloaded-up-to-9-million-times/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/hS0_oWvBHPU" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/family-of-badnews-malware-in-google-play-downloaded-up-to-9-million-times/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/family-of-badnews-malware-in-google-play-downloaded-up-to-9-million-times/</feedburner:origLink></item>
		<item>
		<title>Former Hostgator employee arrested, charged with rooting 2,700 servers</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/Zv4Nbebb3CM/</link>
		<comments>http://arstechnica.com/security/2013/04/former-employee-arrested-charged-with-rooting-2700-hostgator-servers/#comments</comments>
		<pubDate>Fri, 19 Apr 2013 16:51:34 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[hostgator]]></category>
		<category><![CDATA[network security]]></category>
		<category><![CDATA[web hosting]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=249533</guid>
		<description><![CDATA[Prosecutors: Backdoor and digital key gave him near unfettered access.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/server-keys.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							Aurich Lawson				</div>
	  </div>
  </div>
 <p>A former employee of Hostgator has been arrested and charged with installing a backdoor that gave him almost unfettered control over more than 2,700 servers belonging to the widely used Web hosting provider.</p>
<p>Eric Gunnar Gisse, 29, of San Antonio, Texas, was charged with felony breach of computer security by the district attorney's office of Harris County in Texas, according to court documents. He worked as a medium-level administrator from September 2011 until he was terminated on February 15, 2012, according to prosecutors and a company executive. A day after his dismissal, Hostgator officials discovered a backdoor application that allowed Gisse to log in to servers from remote locations, including a computer located at the Hetzner Data Center in Nuremberg, Germany. He took pains to disguise his malware as a widely used Unix administration tool to prevent his superiors from discovering the backdoor process, prosecutors said.</p>
<p>"The process was named 'pcre', a common system file, in order to disguise the true purpose of the process which would grant an attacker unauthorized access into Hostgator's computer network," a Houston Police Department investigator and the document's "affiant," Gordon M. Garrett, wrote in an affidavit. "Complainant told affiant he searched Hostgator's computer network and found the unauthorized 'pcre' process installed on 2723 different Hostgator servers within the computer network."</p>
</div><p><a href="http://arstechnica.com/security/2013/04/former-employee-arrested-charged-with-rooting-2700-hostgator-servers/#p3">Read 7 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/former-employee-arrested-charged-with-rooting-2700-hostgator-servers/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/Zv4Nbebb3CM" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/former-employee-arrested-charged-with-rooting-2700-hostgator-servers/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/former-employee-arrested-charged-with-rooting-2700-hostgator-servers/</feedburner:origLink></item>
		<item>
		<title>Apple remembers where you wanted to get drunk for up to 2 years</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/HF6MhKgo4ss/</link>
		<comments>http://arstechnica.com/apple/2013/04/apple-remembers-where-you-wanted-to-get-drunk-for-up-to-2-years/#comments</comments>
		<pubDate>Fri, 19 Apr 2013 15:30:17 +0000</pubDate>
		<dc:creator>Jacqui Cheng</dc:creator>
				<category><![CDATA[Infinite Loop]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[apple]]></category>
		<category><![CDATA[cloud]]></category>
		<category><![CDATA[data]]></category>
		<category><![CDATA[data retention]]></category>
		<category><![CDATA[privacy]]></category>
		<category><![CDATA[Siri]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=249423</guid>
		<description><![CDATA[Apple keeps a record of Siri queries but says it anonymizes the data.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/siri_getdrunk.png"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">Apple probably still has this query of mine from 2011 saved somewhere in the cloud.</div>
	
	  </div>
  </div>
 <p>Remember that time when you asked Siri about the nearest place to <a href="http://arstechnica.com/apple/2012/10/siri-no-hookers-please-were-chinese/">find hookers</a>? Or perhaps the time you wanted to know where to find the best burritos at 3am? Whatever you've been asking Siri since its launch in late 2011 is likely still on record with Apple, as revealed by a report by our friends at <em><a href="http://www.wired.com/wiredenterprise/2013/04/siri-two-years/">Wired</a></em> on Friday. Apple spokesperson Trudy Muller told <em>Wired</em> that Apple stores Siri queries on its servers for "up to two years," though the company says it makes efforts to anonymize the data.</p>
<p>"Apple may keep anonymized Siri data for up to two years," Muller said. "Our customers’ privacy is very important to us."</p>
<p>Why does Apple have your Siri queries on record in the first place? Remember, Siri doesn't just operate locally on your iPhone or iPad—when you ask it a question, your voice query is sent to Apple's servers for processing before the answer—a Google search, an answer from Wolfram alpha, a Yelp result, etc.—is sent back. That's why an Internet connection is required in order to use Siri; if you have no Wi-Fi or cellular signal, you can't use Siri to perform any actions.</p>
</div><p><a href="http://arstechnica.com/apple/2013/04/apple-remembers-where-you-wanted-to-get-drunk-for-up-to-2-years/#p3">Read 4 remaining paragraphs</a> | <a href="http://arstechnica.com/apple/2013/04/apple-remembers-where-you-wanted-to-get-drunk-for-up-to-2-years/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/HF6MhKgo4ss" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/apple/2013/04/apple-remembers-where-you-wanted-to-get-drunk-for-up-to-2-years/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/apple/2013/04/apple-remembers-where-you-wanted-to-get-drunk-for-up-to-2-years/</feedburner:origLink></item>
		<item>
		<title>Oracle takes a leaf out of Microsoft’s book, prioritizes Java security</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/jmp84l5Kj-Q/</link>
		<comments>http://arstechnica.com/security/2013/04/oracle-takes-a-leaf-out-of-microsofts-book-prioritizes-java-security/#comments</comments>
		<pubDate>Fri, 19 Apr 2013 13:00:24 +0000</pubDate>
		<dc:creator>Peter Bright</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[Java]]></category>
		<category><![CDATA[oracle]]></category>
		<category><![CDATA[security]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=249313</guid>
		<description><![CDATA[Java 8 being delayed into the first quarter of 2014.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>The release of Java 8, originally due in September this year, has been pushed back. The new version's headline feature—Project Lambda, which brings anonymous functions to Java—isn't yet finished.</p>
<p>The reason for this delay is, in part, security. Over the past eight months, a large number of <a href="http://arstechnica.com/security/2012/08/oracle-patches-critical-java-bugs/">critical</a> <a href="http://arstechnica.com/security/2013/01/oracle-patches-widespread-java-zero-day-bug-in-just-three-days-that-is/">security flaws</a> have been found and patched. This has damaged Java's reputation, with Apple, for example, reacting by <a href="http://arstechnica.com/apple/2012/10/apple-removes-java-from-all-os-x-web-browsers/">removing the Java plugin</a> from its Safari browser.</p>
<p>In response, Mark Reinhold, chief architect of the Java Platform Group at Oracle, has <a href="http://mreinhold.org/blog/secure-the-train">announced</a> a "renewed focus on security" that will tie up engineering efforts. As a result, Java 8 has now been pushed back until the first quarter of 2014.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/oracle-takes-a-leaf-out-of-microsofts-book-prioritizes-java-security/#p3">Read 3 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/oracle-takes-a-leaf-out-of-microsofts-book-prioritizes-java-security/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/jmp84l5Kj-Q" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/oracle-takes-a-leaf-out-of-microsofts-book-prioritizes-java-security/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/oracle-takes-a-leaf-out-of-microsofts-book-prioritizes-java-security/</feedburner:origLink></item>
		<item>
		<title>Yes, “design flaw” in 1Password is a problem, just not for end users</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/p6YJzwrXgpU/</link>
		<comments>http://arstechnica.com/security/2013/04/yes-design-flaw-in-1password-is-a-problem-just-not-for-end-users/#comments</comments>
		<pubDate>Thu, 18 Apr 2013 16:52:29 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[1password]]></category>
		<category><![CDATA[cryptography]]></category>
		<category><![CDATA[cryptology]]></category>
		<category><![CDATA[encryption]]></category>
		<category><![CDATA[passwords]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=248417</guid>
		<description><![CDATA[It may very well be time for a new and improved hashing function.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p></p>
<div class="image right full"><img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/keyhole.jpg"></div>Over the past 48 hours, Internet security forums have buzzed with news about a newly discovered technique that allows crackers to make an impressive 3 million guesses per second when trying to find the passcode that unlocks the contents of the widely used <a href="https://agilebits.com/onepassword">1Password password manager</a>.
<p>The optimization, devised by the developer of the <a href="http://hashcat.net/oclhashcat-plus/">oclHashcat-plus password cracking tool</a>, achieved guessing speeds that were, depending on whom you are asking, from two to four times faster than expected. Its discovery was surprising, mainly because it relies in part on a subtle design flaw that until now has been overlooked.</p>
<p>Cryptographers disagree about whether the weakness resides in the popular cryptographic hash function folded into 1Password or the specific implementation contained in 1Password. Either way, the designers of 1Password are smart people who do cryptography right, so the flaw has turned heads. And while even a four-fold reduction in the time it takes to exhaust a cracking attack isn't earth-shattering, it's still significant, considering how many people use 1Password to store the keys to their digital kingdoms.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/yes-design-flaw-in-1password-is-a-problem-just-not-for-end-users/#p3">Read 16 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/yes-design-flaw-in-1password-is-a-problem-just-not-for-end-users/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/p6YJzwrXgpU" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/yes-design-flaw-in-1password-is-a-problem-just-not-for-end-users/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/yes-design-flaw-in-1password-is-a-problem-just-not-for-end-users/</feedburner:origLink></item>
		<item>
		<title>Microsoft rolls out standards-compliant two-factor authentication</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/d0YcGoliR1E/</link>
		<comments>http://arstechnica.com/security/2013/04/microsoft-rolls-out-standards-compliant-two-factor-authentication/#comments</comments>
		<pubDate>Wed, 17 Apr 2013 17:06:22 +0000</pubDate>
		<dc:creator>Peter Bright</dc:creator>
				<category><![CDATA[Gear & Gadgets]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[2-factor authentication]]></category>
		<category><![CDATA[microsoft account]]></category>
		<category><![CDATA[windows phone]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247781</guid>
		<description><![CDATA[Microsoft scheme uses the same tech as Google's two-factor authentication.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>Microsoft today <a href="http://blogs.technet.com/b/microsoft_blog/archive/2013/04/17/microsoft-account-gets-more-secure.aspx">announced</a> that it is rolling out optional two-factor authentication to the 700 million or so Microsoft Account users, confirming last week's <a href="http://arstechnica.com/information-technology/2013/04/two-factor-authentication-finally-heading-to-microsoft-accounts/">rumors</a>. The scheme will become available to all users "in the next few days."</p>
<p>It works essentially identically to existing schemes already available for Google accounts. Two-factor authentication augments a password with a one-time code that's delivered either by text message or generated in an authentication app.</p>
<p>Computers that you trust will be allowed to skip the two factors and just use a password, and application-specific passwords can be generated for interoperating with software that doesn't support two-factor authentication.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/microsoft-rolls-out-standards-compliant-two-factor-authentication/#p3">Read 1 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/microsoft-rolls-out-standards-compliant-two-factor-authentication/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/d0YcGoliR1E" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/microsoft-rolls-out-standards-compliant-two-factor-authentication/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/microsoft-rolls-out-standards-compliant-two-factor-authentication/</feedburner:origLink></item>
		<item>
		<title>Fueled by super botnets, DDoS attacks grow meaner and ever-more powerful</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/QTLIjglO7vc/</link>
		<comments>http://arstechnica.com/security/2013/04/fueled-by-super-botnets-ddos-attacks-grow-meaner-and-ever-more-powerful/#comments</comments>
		<pubDate>Wed, 17 Apr 2013 07:00:00 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[botnets]]></category>
		<category><![CDATA[DDoS]]></category>
		<category><![CDATA[distributed denial of service]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247537</guid>
		<description><![CDATA[Average amount of bandwidth used in DDoS attacks spiked eight-fold last quarter.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/water-cannon-640x457.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="https://commons.wikimedia.org/wiki/File:CPE-Protests-ShowerTime!.jpg">Wikipedia</a>
				</div>
	  </div>
  </div>
 <p>Coordinated attacks used to knock websites offline grew meaner and more powerful in the past three months, with an eight-fold increase in the average amount of junk traffic used to take sites down, according to a company that helps customers weather the so-called distributed denial-of-service campaigns.</p>
<p>The average amount of bandwidth used in DDoS attacks mushroomed to an astounding 48.25 gigabits per second in the first quarter, with peaks as high as 130 Gbps, according to Hollywood, Florida-based Prolexic. During the same period last year, bandwidth in the average attack was 6.1 Gbps and in the fourth quarter of last year it was 5.9 Gbps. The average duration of attacks also grew to 34.5 hours, compared with 28.5 hours last year and 32.2 hours during the fourth quarter of 2012. Earlier this month, Prolexic engineers saw an attack that exceeded 160 Gbps, and officials said they wouldn't be surprised if peaks break the 200 Gbps threshold by the end of June.</p>
<p>The spikes are brought on by new attack techniques that Ars <a href="http://arstechnica.com/security/2012/10/ddos-attacks-against-major-us-banks-no-stuxnet/">first chronicled in October</a>. Rather than using compromised PCs in homes and small offices to flood websites with torrents of traffic, attackers are relying on Web servers, which often have orders of magnitude more bandwidth at their disposal. As Ars reported last week, an ongoing attack on servers running the WordPress blogging application is actively seeking new recruits that can also be harnessed to form <a href="http://arstechnica.com/security/2013/04/huge-attack-on-wordpress-sites-could-spawn-never-before-seen-super-botnet/">never-before-seen botnets</a> to bring still more firepower.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/fueled-by-super-botnets-ddos-attacks-grow-meaner-and-ever-more-powerful/#p3">Read 9 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/fueled-by-super-botnets-ddos-attacks-grow-meaner-and-ever-more-powerful/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/QTLIjglO7vc" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/fueled-by-super-botnets-ddos-attacks-grow-meaner-and-ever-more-powerful/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/fueled-by-super-botnets-ddos-attacks-grow-meaner-and-ever-more-powerful/</feedburner:origLink></item>
		<item>
		<title>ACLU asks feds to probe wireless carriers over Android security updates</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/KaB7aa8w908/</link>
		<comments>http://arstechnica.com/security/2013/04/wireless-carriers-deceptive-and-unfair/#comments</comments>
		<pubDate>Wed, 17 Apr 2013 02:01:25 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Gear & Gadgets]]></category>
		<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[android]]></category>
		<category><![CDATA[carriers]]></category>
		<category><![CDATA[malware]]></category>
		<category><![CDATA[security]]></category>
		<category><![CDATA[smartphones]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247371</guid>
		<description><![CDATA["Defective" phones from AT&#038;T, Verizon, Sprint, T-Mobile pose risks, ACLU says.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/android-sharks.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							Aurich Lawson / Thinkstock				</div>
	  </div>
  </div>
 <p>Civil liberties advocates have asked the US Federal Trade Commission to take action against the nation's four major wireless carriers for selling millions of Android smartphones that never, or only rarely, receive updates to patch dangerous security vulnerabilities.</p>
<p>The request for investigation and complaint for injunctive relief was filed Tuesday by the American Civil Liberties Union against AT&amp;T, Verizon Wireless, Sprint Nextel, and T-Mobile USA. The majority of phones that the carriers sell run Google's Android operating system and rarely receive software updates, the 16-page document stated. It went on to allege that the practice violates provisions of the Federal Trade Commission Act barring deceptive and unfair business practices, since the carriers don't disclose that the failure to provide updates in a timely manner puts customers at greater risk of hacking attacks. Among other things, the filing seeks an order allowing customers to terminate contracts that cover a phone that's no longer eligible to receive updates.</p>
<p>"All four of the major wireless carriers consistently fail to provide consumers with available security updates to repair known security vulnerabilities in the software operating on mobile devices," Christopher Soghoian, principal technologist and senior policy analyst for the ACLU, wrote in the document. "The wireless carriers have failed to warn consumers that the smartphones sold to them are defective and that they are running vulnerable operating system and browser software. The delivery of software updates to consumers is not just an industry best practice, but is in fact a basic requirement for companies selling computing devices that they know will be used to store sensitive information, such as intimate photographs, e-mail, instant messages, and online banking credentials."</p>
</div><p><a href="http://arstechnica.com/security/2013/04/wireless-carriers-deceptive-and-unfair/#p3">Read 14 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/wireless-carriers-deceptive-and-unfair/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/KaB7aa8w908" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/wireless-carriers-deceptive-and-unfair/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/wireless-carriers-deceptive-and-unfair/</feedburner:origLink></item>
		<item>
		<title>ColdFusion hack used to steal hosting provider’s customer data</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/SqDlyNS9RoM/</link>
		<comments>http://arstechnica.com/security/2013/04/coldfusion-hack-used-to-steal-hosting-providers-customer-data/#comments</comments>
		<pubDate>Tue, 16 Apr 2013 22:45:57 +0000</pubDate>
		<dc:creator>Sean Gallagher</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[Adobe]]></category>
		<category><![CDATA[breach]]></category>
		<category><![CDATA[ColdFusion]]></category>
		<category><![CDATA[Linode]]></category>
		<category><![CDATA[vulnerability]]></category>
		<category><![CDATA[zero-day]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247267</guid>
		<description><![CDATA[Linode hit by possible zero-day exploit patched by Adobe on April 9.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>A vulnerability in the ColdFusion Web server platform, reported by Adobe less than a week ago, has apparently been in the wild for almost a month and has allowed the hacking of at least one company website, exposing customer data. Yesterday, it was revealed that the virtual server hosting company Linode had been the victim of a multi-day breach that allowed hackers to gain access to customer records.</p>
<p>The breach was made possible by a vulnerability in Adobe's ColdFusion server platform that could, according to Adobe, "be exploited to impersonate an authenticated user." A patch had been issued for the vulnerability on <a href="http://www.adobe.com/support/security/bulletins/apsb13-10.html">April 9</a> and was rated as priority "2" and "important." Those ratings placed it at a step down from the most critical, indicating that there were no known exploits at the time the patch was issued but that data was at risk. Adobe credited "an anonymous security researcher," with discovering the vulnerability.</p>
<p>But according to <a href="http://turtle.dereferenced.org/~nenolod/linode/linode-abridged.txt">IRC conversation</a> including one of the alleged hackers of the site, Linode's site had been compromised for weeks before its discovery. That revelation leaves open the possibility that other ColdFusion sites have been compromised as hackers sought out targets to use the exploit on.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/coldfusion-hack-used-to-steal-hosting-providers-customer-data/#p3">Read 5 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/coldfusion-hack-used-to-steal-hosting-providers-customer-data/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/SqDlyNS9RoM" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/coldfusion-hack-used-to-steal-hosting-providers-customer-data/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/coldfusion-hack-used-to-steal-hosting-providers-customer-data/</feedburner:origLink></item>
		<item>
		<title>“Syrian Electronic Army” hacks NPR publishing system, edits articles</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/m0ELa49tXXk/</link>
		<comments>http://arstechnica.com/information-technology/2013/04/syrian-electronic-army-hacks-npr-publishing-system-edits-articles/#comments</comments>
		<pubDate>Tue, 16 Apr 2013 16:03:11 +0000</pubDate>
		<dc:creator>Jon Brodkin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[hack]]></category>
		<category><![CDATA[NPR]]></category>
		<category><![CDATA[Syria]]></category>
		<category><![CDATA[syrian electronic army]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=246713</guid>
		<description><![CDATA[Headlines changed to "Syrian Electronic Army was here."
]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>NPR's Web publishing system and several of the news agency's Twitter accounts were hacked yesterday by a group supportive of the Syrian government that calls itself the "Syrian Electronic Army."</p>
<p>"Late Monday evening, several stories on the NPR website were defaced with headlines and text that said 'Syrian Electronic Army Was Here,'" an NPR statement published in a <a href="http://NPR.org">NPR.org</a><a href="http://www.npr.org/blogs/thetwo-way/2013/04/16/177421655/npr-org-hacked-syrian-electronic-army-takes-credit">news story</a> on the incident said. "Some of these stories were distributed to and appeared on NPR Member Station websites. We have made the necessary corrections to those stories on NPR.org and are continuing to work with our Member Stations. Similar statements were posted on several NPR Twitter accounts. Those Twitter accounts have been addressed. We are closely monitoring the situation."</p>
<p>Sophos's Naked Security blog published a <a href="http://nakedsecurity.sophos.com/2013/04/16/syrian-electronic-army-npr/">summary of the hack</a>, including a screenshot of a Google search showing some of the headlines edited by the Syrian Electronic Army:</p>
</div><p><a href="http://arstechnica.com/information-technology/2013/04/syrian-electronic-army-hacks-npr-publishing-system-edits-articles/#p3">Read 3 remaining paragraphs</a> | <a href="http://arstechnica.com/information-technology/2013/04/syrian-electronic-army-hacks-npr-publishing-system-edits-articles/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/m0ELa49tXXk" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/information-technology/2013/04/syrian-electronic-army-hacks-npr-publishing-system-edits-articles/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/information-technology/2013/04/syrian-electronic-army-hacks-npr-publishing-system-edits-articles/</feedburner:origLink></item>
		<item>
		<title>New security protection, fixes for 39 exploitable bugs coming to Java</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/h-rhBnccfP0/</link>
		<comments>http://arstechnica.com/security/2013/04/new-security-protection-fixes-for-39-exploitable-bugs-coming-to-java/#comments</comments>
		<pubDate>Mon, 15 Apr 2013 20:30:40 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[computer security]]></category>
		<category><![CDATA[exploit]]></category>
		<category><![CDATA[Java]]></category>
		<category><![CDATA[malware]]></category>
		<category><![CDATA[oracle]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=246297</guid>
		<description><![CDATA[Latest version is designed to help users block <em>more</em> Java exploits on websites.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div style="float:right top">
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/unsigned_cert.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">A dialog box presented by Java when it encounters an application that isn't signed by a digital certificate.</div>
	
			<div class="caption-credit">
							<a rel="nofollow" href="https://www.java.com/en/download/help/appsecuritydialogs.xml">Java.com</a>
				</div>
	  </div>
  </div>
 <p>Oracle plans to release an update for the widely exploited Java browser plugin. The update fixes 39 critical vulnerabilities and introduces changes designed to make it harder to carry out drive-by attacks on end-user computers.</p>
<p>The update scheduled for Tuesday comes as the security of Java is reaching near-crisis levels. Throughout the past year, a series of attacks hosted on popular websites has been used to surreptitiously install malware on unwitting users' machines. The security flaws have been used to infect employees of <a href="http://arstechnica.com/security/2013/02/dev-site-behind-apple-facebook-hacks-didnt-know-it-was-booby-trapped/">Facebook and Apple</a> in targeted attacks intended to penetrate those companies. The vulnerabilities have also been exploited to hijack computers of home and business users. More than once, attackers have <a href="http://arstechnica.com/security/2013/03/another-java-zero-day-exploit-in-the-wild-actively-attacking-targets/">exploited one previously undocumented bug</a> within days or weeks of <a href="http://arstechnica.com/security/2013/02/javas-latest-security-problems-new-flaw-identified-old-one-attacked/">patching a previous "zero-day,"</a> as such vulnerabilities are known, creating a string of attacks on the latest version of the widely used plugin.</p>
<p>In all, Java 7 Update 21 will fix at least 42 security bugs, Oracle said in a <a href="http://www.oracle.com/technetwork/topics/security/javacpuapr2013-1928497.html">pre-release announcement</a>. The post went on to say that "39 of those vulnerabilities may be remotely exploitable without authentication, i.e., may be exploited over a network without the need for a username and password." The advisory didn't specify or describe the holes that will be patched. Security Exploration, a Poland-based security company that has discovered dozens of "security issues" in Java, has a running list of them <a href="http://www.security-explorations.com/en/SE-2012-01-status.html">here</a>.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/new-security-protection-fixes-for-39-exploitable-bugs-coming-to-java/#p3">Read 5 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/new-security-protection-fixes-for-39-exploitable-bugs-coming-to-java/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/h-rhBnccfP0" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/new-security-protection-fixes-for-39-exploitable-bugs-coming-to-java/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/new-security-protection-fixes-for-39-exploitable-bugs-coming-to-java/</feedburner:origLink></item>
		<item>
		<title>Huge attack on WordPress sites could spawn never-before-seen super botnet</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/eq3OrFwFTBo/</link>
		<comments>http://arstechnica.com/security/2013/04/huge-attack-on-wordpress-sites-could-spawn-never-before-seen-super-botnet/#comments</comments>
		<pubDate>Sat, 13 Apr 2013 01:10:52 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[wordpress]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=245759</guid>
		<description><![CDATA[Ongoing attack from >90,000 computers is creating a strain on Web hosts, too.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/wp_bruteforce-640x455.png"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://blog.cloudflare.com/patching-the-internet-fixing-the-wordpress-br">CloudFlare</a>
				</div>
	  </div>
  </div>
 <p>Security analysts have detected an ongoing attack that uses a huge number of computers from across the Internet to commandeer servers that run the WordPress blogging application.</p>
<p>The unknown people behind the highly distributed attack are using more than 90,000 IP addresses to brute-force crack administrative credentials of vulnerable WordPress systems, researchers from at least three Web hosting services reported. At least one company warned that the attackers may be in the process of building a "botnet" of infected computers that's vastly stronger and more destructive than those available today. That's because the servers have bandwidth connections that are typically tens, hundreds, or even thousands of times faster than botnets made of infected machines in homes and small businesses.</p>
<p>"These larger machines can cause much more damage in DDoS [distributed denial-of-service] attacks because the servers have large network connections and are capable of generating significant amounts of traffic," Matthew Prince, CEO of content delivery network CloudFlare, wrote in a <a href="http://blog.cloudflare.com/patching-the-internet-fixing-the-wordpress-br">blog post</a> describing the attacks.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/huge-attack-on-wordpress-sites-could-spawn-never-before-seen-super-botnet/#p3">Read 10 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/huge-attack-on-wordpress-sites-could-spawn-never-before-seen-super-botnet/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/eq3OrFwFTBo" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/huge-attack-on-wordpress-sites-could-spawn-never-before-seen-super-botnet/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/huge-attack-on-wordpress-sites-could-spawn-never-before-seen-super-botnet/</feedburner:origLink></item>
		<item>
		<title>Microsoft tells Windows 7 users to uninstall faulty security update (Updated)</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/bfKsueVgUos/</link>
		<comments>http://arstechnica.com/security/2013/04/microsoft-tells-windows-7-users-to-uninstall-faulty-security-update/#comments</comments>
		<pubDate>Fri, 12 Apr 2013 18:50:58 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[compter security]]></category>
		<category><![CDATA[malware]]></category>
		<category><![CDATA[microsoft]]></category>
		<category><![CDATA[windows 7]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=245117</guid>
		<description><![CDATA[Patch causes some machines to become unbootable, company warns.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>Microsoft has pulled a Windows 7 security update released as part of this month's <a href="https://blogs.technet.com/b/srd/archive/2013/04/09/assessing-risk-for-the-april-2013-security-updates.aspx?Redirected=true">Patch Tuesday</a> after discovering it caused some machines to become unbootable.</p>
<p><a href="http://support.microsoft.com/kb/2823324">Update 2823324</a>, which was included in the <a href="http://technet.microsoft.com/en-us/security/bulletin/ms13-036">MS13-036 bulletin</a>, fixed a "moderate-level vulnerability" that requires an attacker to have physical computer access to be able to exploit a targeted computer, Dustin Childs, a group manager in the Microsoft Trustworthy Computing group, wrote in a <a href="https://blogs.technet.com/b/msrc/archive/2013/04/11/kb2839011-released-to-address-security-bulletin-update-issue.aspx?Redirected=true">blog post</a> published Thursday evening. The company has now pulled it from the bulletin and is advising at least some Windows users who have installed it to uninstall the update following the <a href="http://support.microsoft.com/kb/2839011">guidance here</a>. MS130-26 was one of nine bulletins released on Monday to fix 13 separate vulnerabilities.</p>
<p>"We’ve determined that the update, when paired with certain third-party software, can cause system errors," Childs wrote. "As a precaution, we stopped pushing 2823324 as an update when we began investigating the error reports, and have since removed it from the download center."</p>
</div><p><a href="http://arstechnica.com/security/2013/04/microsoft-tells-windows-7-users-to-uninstall-faulty-security-update/#p3">Read 3 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/microsoft-tells-windows-7-users-to-uninstall-faulty-security-update/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/bfKsueVgUos" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/microsoft-tells-windows-7-users-to-uninstall-faulty-security-update/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/microsoft-tells-windows-7-users-to-uninstall-faulty-security-update/</feedburner:origLink></item>
		<item>
		<title>A beginner’s guide to building botnets—with little assembly required</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/2e_zUqhObQ4/</link>
		<comments>http://arstechnica.com/security/2013/04/a-beginners-guide-to-building-botnets-with-little-assembly-required/#comments</comments>
		<pubDate>Fri, 12 Apr 2013 01:00:08 +0000</pubDate>
		<dc:creator>Sean Gallagher</dc:creator>
				<category><![CDATA[Features]]></category>
		<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[botnets]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=241731</guid>
		<description><![CDATA[For a few hundred dollars, you can get tools and 24/7 support for Internet crime.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/Botnet-Vending-Machine-640x398.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.flickr.com/photos/m-i-k-e/6496110811/">Original photo by Michael Kappel / Remixed by Aurich Lawson</a>
				</div>
	  </div>
  </div>
 <p>Have a plan to steal millions from banks and their customers but can't write a line of code? Want to get rich quick off advertising click fraud but "quick" doesn't include time to learn how to do it? No problem. Everything you need to start a life of cybercrime is just a few clicks (and many more dollars) away.</p>
<p>Building successful malware is an expensive business. It involves putting together teams of developers, coordinating an army of fraudsters to convert ill-gotten gains to hard currency without pointing a digital arrow right back to you. So the biggest names in financial botnets—Zeus, Carberp, Citadel, and SpyEye, to name a few—have all at one point or another decided to shift gears from fraud rings to crimeware vendors, selling their wares to whoever can afford them.</p>
<p>In the process, these big botnet platforms have created a whole ecosystem of software and services in an underground market catering to criminals without the skills to build it themselves. As a result, the tools and techniques used by last years' big professional bank fraud operations, such as the "<a href="http://arstechnica.com/tech-policy/2012/06/sophisticated-bank-fraud-attempted-to-steal-at-least-78-million/">Operation High Roller</a>" botnet that netted over $70 million last summer, are available off-the-shelf on the Internet. They even come with full technical support to help you get up and running.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/a-beginners-guide-to-building-botnets-with-little-assembly-required/#p3">Read 63 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/a-beginners-guide-to-building-botnets-with-little-assembly-required/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/2e_zUqhObQ4" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/a-beginners-guide-to-building-botnets-with-little-assembly-required/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/a-beginners-guide-to-building-botnets-with-little-assembly-required/</feedburner:origLink></item>
		<item>
		<title>More than 30 MMORPG companies targeted in ongoing malware attack</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/YitfMqtKYrI/</link>
		<comments>http://arstechnica.com/security/2013/04/more-then-30-mmorpg-companies-targeted-in-ongoing-malware-attack/#comments</comments>
		<pubDate>Fri, 12 Apr 2013 00:00:37 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Opposable Thumbs]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[MMORPG]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=244617</guid>
		<description><![CDATA[In at least two cases, malware was planted on update servers and spread to fans.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>Researchers have uncovered an ongoing cyberespionage campaign targeting more than 30 online video game companies over the past four years.</p>
<p>The companies infected by the malware primarily market so-called massively multiplayer online role-playing games. They're mostly located in South East Asia, but are also in the US, Germany, Japan, China, Russia, Brazil, Peru, and Belarus, according to a <a href="https://www.securelist.com/en/blog/855/Winnti_FAQ_More_than_just_a_game">release published Thursday</a> by researchers from antivirus provider Kaspersky Lab. The attackers work from computers with Chinese and Korean language configurations. They used their unauthorized access to obtain digital certificates that were later exploited in malware campaigns targeting other industries and political activists.</p>
<p>So far, there's no evidence that customers of the infected game companies were targeted, although in at least one case, malicious code was accidentally installed on gamers' computers by one of the infected victim companies. Kaspersky said there was another case of end users being infected by the malware, which is known as "Winnti." The company didn't rule out the possibility that players could be hit in the future, potentially as a result of collateral damage.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/more-then-30-mmorpg-companies-targeted-in-ongoing-malware-attack/#p3">Read 4 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/more-then-30-mmorpg-companies-targeted-in-ongoing-malware-attack/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/YitfMqtKYrI" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/more-then-30-mmorpg-companies-targeted-in-ongoing-malware-attack/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/more-then-30-mmorpg-companies-targeted-in-ongoing-malware-attack/</feedburner:origLink></item>
		<item>
		<title>US adds Russian supercomputer maker to list of nuclear threats</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/NtnayBcYS7U/</link>
		<comments>http://arstechnica.com/information-technology/2013/04/us-adds-russian-supercomputer-maker-to-list-of-nuclear-threats/#comments</comments>
		<pubDate>Thu, 11 Apr 2013 20:15:24 +0000</pubDate>
		<dc:creator>Jon Brodkin</dc:creator>
				<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[nuclear]]></category>
		<category><![CDATA[supercomputer]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=244327</guid>
		<description><![CDATA[Maker of world's 26th fastest supercomputer faces heavy trade restrictions.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/t-platform-cluster-640x403.png"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">A T-Platforms supercomputer.</div>
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.t-platforms.com/solutions.html">T-Platforms</a>
				</div>
	  </div>
  </div>
 <p>Six months ago, a company called T-Platforms triumphantly <a href="http://www.t-platforms.com/about-company/press-releases/395-the-oganov-lab.html">announced</a> the "First Delivery of [a] Russian Supercomputer to [the] US."</p>
<p>The US government has since added T-Platforms to a <a href="http://www.bis.doc.gov/policiesandregulations/ear/744_supp4.pdf">list of entities</a> that are "acting contrary to the national security or foreign policy interests of the United States" by having involvement with nuclear research. Specifically, T-Platforms' operations in Russia, Germany, and Taiwan were added to the Export Administration Regulations (EAR) Entity List by representatives of the US Departments of Commerce, State, Defense, and Energy. This will make it difficult for T-Platforms to do business with US companies, although it isn't an outright ban.</p>
<p>"The Entity List notifies the public about entities that have engaged in activities that could result in an increased risk of the diversion of exported, reexported, or transferred (in-country) items to weapons of mass destruction (WMD) programs," the Department of Commerce's Bureau of Industry and Security said in its <a href="http://www.gpo.gov/fdsys/pkg/FR-2013-03-08/html/2013-05387.htm">notice</a> that T-Platforms is now on the list. "Since its initial publication, grounds for inclusion on the Entity List have expanded to activities sanctioned by the State Department and activities contrary to U.S. National security or foreign policy interests, including terrorism and export control violations involving abuse of human rights."</p>
</div><p><a href="http://arstechnica.com/information-technology/2013/04/us-adds-russian-supercomputer-maker-to-list-of-nuclear-threats/#p3">Read 10 remaining paragraphs</a> | <a href="http://arstechnica.com/information-technology/2013/04/us-adds-russian-supercomputer-maker-to-list-of-nuclear-threats/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/NtnayBcYS7U" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/information-technology/2013/04/us-adds-russian-supercomputer-maker-to-list-of-nuclear-threats/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/information-technology/2013/04/us-adds-russian-supercomputer-maker-to-list-of-nuclear-threats/</feedburner:origLink></item>
		<item>
		<title>Hacking commercial aircraft with an Android App (some conditions apply)</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/security/~3/U9LdvbWC1mI/</link>
		<comments>http://arstechnica.com/security/2013/04/hacking-commercial-aircraft-with-an-android-app-some-conditions-apply/#comments</comments>
		<pubDate>Thu, 11 Apr 2013 19:32:00 +0000</pubDate>
		<dc:creator>Dan Goodin</dc:creator>
				<category><![CDATA[Risk Assessment]]></category>
		<category><![CDATA[Technology Lab]]></category>
		<category><![CDATA[hardware hacking]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=244325</guid>
		<description><![CDATA[Comms system used in the aviation industry contains no security, researcher says.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>As if <a href="http://arstechnica.com/security/2012/12/how-to-bring-down-mission-critical-gps-networks-with-2500/">inexpensive attacks on mission-critical global positioning systems</a> weren't enough, a researcher said he's developed an Android app that could redirect airplanes in mid-flight.</p>
<p>The frightening scenario was presented on Wednesday at the Hack in the Box security conference in Amsterdam. It's made possible by security weaknesses in the protocol used to send data to commercial planes and in flight-management software built by companies including Honeywell, Thales, and Rockwell Collins, <a href="http://www.forbes.com/sites/andygreenberg/2013/04/10/researcher-says-hes-found-hackable-flaws-in-airplanes-navigation-systems/">Forbes reports</a>. Vulnerable systems include the Aircraft Communications Addressing and Report System used for exchanging text messages between planes and ground stations using VHF radio or satellite signals. It has "virtually no authentication features to prevent spoofed commands."</p>
<p>Using a custom-developed Android app dubbed PlaneSploit, researcher Hugo Tesa of N.Runs showed how a virtual plane in a laboratory could be redirected. Because there's no means to cryptographically authenticate communications sent over ACARS, pilots have no way to confirm if messages they receive in the cockpit are valid. Malformed messages can then be used to trigger vulnerabilities, Tesa told Forbes.</p>
</div><p><a href="http://arstechnica.com/security/2013/04/hacking-commercial-aircraft-with-an-android-app-some-conditions-apply/#p3">Read 3 remaining paragraphs</a> | <a href="http://arstechnica.com/security/2013/04/hacking-commercial-aircraft-with-an-android-app-some-conditions-apply/?comments=1">Comments</a></p><img src="http://feeds.feedburner.com/~r/arstechnica/security/~4/U9LdvbWC1mI" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/security/2013/04/hacking-commercial-aircraft-with-an-android-app-some-conditions-apply/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/security/2013/04/hacking-commercial-aircraft-with-an-android-app-some-conditions-apply/</feedburner:origLink></item>
	</channel>
</rss>
`
