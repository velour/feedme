package feedme

var feeds = []string{
	feed0,
	feed1,
}

const feed0 = `
<?xml version="1.0" encoding="UTF-8"?>
<rss xmlns:atom="http://www.w3.org/2005/Atom" version="2.0" xmlns:posterous="http://posterous.com/help/rss/1.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:media="http://search.yahoo.com/mrss/">
<channel>
<title>CloudFlare blog</title>
<link>http://blog.cloudflare.com</link>
<description>Musings from the CloudFlare team</description>
<generator>posterous.com</generator>
<link type="application/json" href="http://posterous.com/api/sup_update#8ab5e4799" xmlns="http://www.w3.org/2005/Atom" rel="http://api.friendfeed.com/2008/03#sup"/>
<atom:link href="http://blog.cloudflare.com/rss.xml" rel="self"/>
<atom:link href="http://posterous.superfeedr.com" rel="hub"/>
<item>
<pubDate>Thu, 11 Apr 2013 16:23:00 -0700</pubDate>
<title>Patching the Internet in Realtime: Fixing the Current WordPress Brute Force Attack</title>
<link>http://blog.cloudflare.com/patching-the-internet-fixing-the-wordpress-br</link>
<guid>http://blog.cloudflare.com/patching-the-internet-fixing-the-wordpress-br</guid>
<description>
<![CDATA[<p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-04-11/ppjEifHaJDEqyzfafoiecrmypckcfjdtDdsvhHqzBBtxGqjtkwDCEAzrzJnG/wp_bruteforce_opt1.png.scaled1000.png"><img alt="Wp_bruteforce_opt1" height="356" src="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-04-11/ppjEifHaJDEqyzfafoiecrmypckcfjdtDdsvhHqzBBtxGqjtkwDCEAzrzJnG/wp_bruteforce_opt1.png.scaled500.png" width="500"/></a>
</div>
</p>
<p>There is currently a significant attack being launched at a large number of WordPress blogs across the Internet. The attacker is brute force attacking the WordPress administrative portals, using the username "admin" and trying thousands of passwords. It appears a botnet is being used to launch the attack and more than tens of thousands of unique IP addresses have been recorded attempting to hack WordPress installs.</p>
<p>One of the concerns of an attack like this is that the attacker is using a relatively weak botnet of home PCs in order to build a much larger botnet of beefy servers in preparation for a future attack. These larger machines can cause much more damage in DDoS attacks because the servers have large network connections and are capable of generating significant amounts of traffic. This is a similar tactic that was used to <a href="http://www.informationweek.com/security/attacks/bank-attackers-used-php-websites-as-laun/240144413" target="_blank">build the so-called itsoknoproblembro/Brobot botnet </a>which, in the Fall of 2012, was behind the large attacks on US financial institutions.</p>
<p><strong>Patching the Internet</strong></p>
<p>We just pushed a rule out through CloudFlare's WAF that detects the signature of the attack and stops it. Rather than limiting this to only paying customers, CloudFlare is rolling it out the fix to all our customers automatically, including customers on our free plan. If you are a WordPress user and you are using CloudFlare, you are now protected from this latest brute force attack.</p>
<p>Because CloudFlare sits in front of a significant portion of web requests we have the opportunity to, literally, patch Internet vulnerabilities in realtime. We will be providing information about the attack back to partners who are interested in hardening their internal defenses for customers who are not yet on CloudFlare.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Internet_patch" height="250" src="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-04-11/efzFwacolnIvxhspvEBuApxzvxeCfbnjlJevaJJBzFDcnlqFphBwygghqBEn/internet_patch.png.scaled500.png" width="250"/>
</div>
</p>
<p>If you are running a WordPress blog and want to ensure you are protected from this attack, you can <a href="https://www.cloudflare.com/sign-up" target="_blank">sign up for CloudFlare's free plan</a> and the protection is automatic. We'll continue to monitor the details of the attack and publish details about what we learn.</p>
</p>
<p><a href="http://blog.cloudflare.com/patching-the-internet-fixing-the-wordpress-br">Permalink</a>
| <a href="http://blog.cloudflare.com/patching-the-internet-fixing-the-wordpress-br#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>
</p>]]>
</description>
<posterous:author>
<posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
<posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
<posterous:firstName>Matthew</posterous:firstName>
<posterous:lastName>Prince</posterous:lastName>
<posterous:nickName>eastdakota</posterous:nickName>
<posterous:displayName>Matthew Prince</posterous:displayName>
</posterous:author>
<media:content type="image/png" height="250" width="250" url="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-04-11/efzFwacolnIvxhspvEBuApxzvxeCfbnjlJevaJJBzFDcnlqFphBwygghqBEn/internet_patch.png">
<media:thumbnail height="250" width="250" url="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-04-11/efzFwacolnIvxhspvEBuApxzvxeCfbnjlJevaJJBzFDcnlqFphBwygghqBEn/internet_patch.png.scaled500.png"/>
</media:content>
<media:content type="image/png" height="504" width="708" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-04-11/ppjEifHaJDEqyzfafoiecrmypckcfjdtDdsvhHqzBBtxGqjtkwDCEAzrzJnG/wp_bruteforce_opt1.png">
<media:thumbnail height="356" width="500" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-04-11/ppjEifHaJDEqyzfafoiecrmypckcfjdtDdsvhHqzBBtxGqjtkwDCEAzrzJnG/wp_bruteforce_opt1.png.scaled500.png"/>
</media:content>
</item>
<item>
<pubDate>Wed, 10 Apr 2013 18:12:49 -0700</pubDate>
<title>Continuing the Conversations from Parallels Summit 2013</title>
<link>http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a-15903</link>
<guid>http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a-15903</guid>
<description>
<![CDATA[<p>
<p class="p1"><a href="http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a" target="_blank">As we wrote about before</a>, we attended Parallels Summit in February where we hosted &ldquo;Conversations with CloudFlare&rdquo; - live video interviews with industry experts.</p>
<p class="p1">Below are highlights and video clips from five of those conversations. Tune in to hear the latest in mobile services, how hosting providers are expanding across the world, and what these companies are looking forward to in 2013.</p>
<p class="p2">&nbsp;</p>
<p class="p3"><strong><span class="s1"><a href="https://twitter.com/enoss">Elliot Noss</a></span>, CEO of TuCows<br/></strong><em>When it comes to mobile, TuCows is giving AT&amp;T and Verizon a run for their money with <a href="https://ting.com/"><span class="s1">Ting</span></a>. Elliot explains the benefits of Ting, what's new with their mobile service and why wifi on your phone is all the rage...&nbsp;</em></p>
<p class="p4"><iframe src="http://player.vimeo.com/video/61743344?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p class="p5">&nbsp;</p>
<p class="p4"><strong><span class="s1"><a href="https://twitter.com/Hostnet">Merjin de Brabander</a></span>, Business Manager at Hostnet</strong><br/> <em>For 2013, Hostnet's application and POA program is really starting to take off. Listen in as Merjin discusses how they've grown and what new services they're bringing to their customers in 2013...</em></p>
<p class="p4"><iframe src="http://player.vimeo.com/video/61743346?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p class="p5">&nbsp;</p>
<p class="p4"><strong><span class="s1"><a href="https://twitter.com/ViUX">JT Smith</a></span>, Owner of ViUX</strong><br/> <em>Providing an arrange of hosting services, ViUX is improving performance and reliability for their customers and&nbsp;considers themselves a complete Parallels shop.&nbsp;One area they pride themselves on? Customer service...&nbsp;</em></p>
<p class="p6"><iframe src="http://player.vimeo.com/video/61316383?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p class="p5">&nbsp;</p>
<p class="p4"><strong><span class="s1"><a href="https://twitter.com/LimestoneInc">Kris Anderson</a></span>, Director of Support at Limestone Networks</strong><br/> <em>Limestone Networks is in the business of mission critical data. Their infrastructure supports some of the most reliable data centers around. Their support motto? Simple. Solid. Superior.&nbsp;</em></p>
<p class="p4"><iframe src="http://player.vimeo.com/video/61316379?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p class="p5">&nbsp;</p>
<p class="p4"><strong><span class="s1"><a href="https://twitter.com/atomia">Magnus Hult</a></span>, Executive Vice President at Atomia</strong><br/> <em>Atomia is making it easier than ever to run a business. They focus on the automation and billing platform for hosting providers so businesses can focus on what they do best. Hear what Atomia is working on for 2013...</em></p>
<p class="p4"><iframe src="http://player.vimeo.com/video/61316382?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p>&nbsp;</p>
</p>
<p><a href="http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a-15903">Permalink</a>
| <a href="http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a-15903#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>
</p>]]>
</description>
<posterous:author>
<posterous:userImage>http://files.posterous.com/user_profile_pics/1122929/CloudFlare4-16.jpeg</posterous:userImage>
<posterous:profileUrl>http://posterous.com/users/he6rIlp3ceWH8</posterous:profileUrl>
<posterous:firstName>Kristin</posterous:firstName>
<posterous:lastName>Tarr</posterous:lastName>
<posterous:nickName>ktarr</posterous:nickName>
<posterous:displayName>Kristin Tarr</posterous:displayName>
</posterous:author>
</item>
<item>
<pubDate>Mon, 08 Apr 2013 02:18:29 -0700</pubDate>
<title>How the CloudFlare Team Got Into Bondage (It's Not What You Think)</title>
<link>http://blog.cloudflare.com/how-the-cloudflare-team-got-into-bondage-its</link>
<guid>http://blog.cloudflare.com/how-the-cloudflare-team-got-into-bondage-its</guid>
<description>
<![CDATA[<p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-04-07/DhwAldFqkAruHepvqiHehettJJCoyJulEdipGnorkhxuyFgdirahmoqnwqmE/cat5-o-nine-tails.png.scaled1000.png"><img alt="Cat5-o-nine-tails" height="336" src="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-04-07/DhwAldFqkAruHepvqiHehettJJCoyJulEdipGnorkhxuyFgdirahmoqnwqmE/cat5-o-nine-tails.png.scaled500.png" width="500"/></a>
</div>
<span style="font-size: xx-small;">(Image courtesy of ferelswirl)</span></p>
<p>At CloudFlare, we're always looking for ways to eliminate bottlenecks. We're only able to deal with the very large amount of traffic that we handle, especially during large denial of service attacks, because we've built a network that can efficiently handle an extremely high volume of network requests. This post is about the nitty gritty of port bonding, one of the technologies we use, and how it allows us to get the maximum possible network throughput out of our servers.</p>
<p><strong>Generation Three</strong></p>
<p>A rack of equipment in CloudFlare's network has three core components: routers, switches, and servers. We own and install all our own equipment because it's impossible to have the flexibility and efficiency you need to do what we do running on someone else's gear. Over time, we've adjusted the specs of the gear we use based on the needs of our network and what we are able to cost effectively source from vendors.</p>
<p>Most of the equipment in our network today is based on our Generation 3 (G3) spec, which we deployed throughout 2012. Focusing just on the network connectivity for our G3 gear, our routers have multiple 10Gbps ports which connect out to the Internet as well as in to our switches. Our switches have a handful of 10Gbps ports that we use to connect to our routers and then 48 1Gbps ports that connect to the servers. Finally, our servers have 6 1Gbps ports, two on the motherboard (using Intel's chipset) and four on an Intel PCI network card. (There's an additional IPMI management port as well, but it doesn't figure into this discussion.)</p>
<p><div class='p_embed p_image_embed'>
<img alt="Cloudflare_servers" height="640" src="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-04-07/rifaGwFyzhhpGwlvDBesrnwayGartdkIyGpillzzuiakfpDidCHnyvakgADy/cloudflare_servers.jpg.scaled500.jpg" width="480"/>
</div>
</p>
<p>To get high levels of utilization and keep our server spec consistent and flexible, each of the servers in our network can perform any of the key CloudFlare functions: DNS, front-line, caching, and logging. Cache, for example, is spread across multiple machines in a facility. This means if we add another drive to one of the servers in a data center, then the total available storage space for the cache increases for all the servers in that data center. What's good about this is that, as we need to, we can add more servers and linearly scale capacity across storage, CPU, and, in some applications, RAM. The challenge is that in order to pull this off there needs to be a significant amount of communication between servers across our local area network (LAN).</p>
<p>When we originally started deploying our G3 servers in early 2012, we treated each 1Gbps port on the switches and routers discretely. While each server could, in theory, handle 6Gbps of traffic, each port could only handle 1Gbps. Usually this was no big deal because we load balanced customers across multiple servers in multiple data centers so on no individual server port was a customer likely to burst over 1Gbps. However, we found that, from time to time, when a customer would come under attack, traffic to individual machines could exceed 1Gbps and overwhelm a port.</p>
<p><strong>When A Problem Comes Along...</strong></p>
<p>The goal of a denial of service attack is to find a bottleneck and then send enough garbage requests to fill it up and prevent legitimate requests from getting through. At the same time, our goal when mitigating such an attack is first to ensure the attack doesn't harm other customers and then to stop the attack from hurting the actual target.</p>
<p><img class="posterous_plugin_object posterous_plugin_object_image" src="http://24.media.tumblr.com/tumblr_m54e8tjzdQ1qfj10wo1_500.gif" alt="Devo Whip It"/></p>
<p>For the most part, the biggest attacks by volume we see are Layer 3 attacks. In these, packets are stopped at the edge of our network and never reach our server infrastructure. As the <a href="http://blog.cloudflare.com/the-ddos-that-almost-broke-the-internet" target="_blank">very large attack against Spamhaus</a> showed, we have a significant amount of network capacity at our edge and are therefore able to stop these Layer 3 attacks very effectively.</p>
<p>While the big Layer 3 attacks get the most attention, an attack doesn't need to be so large if it can affect another, narrower bottleneck. For example, switches and routers are largely blind to Layer 7 attacks, meaning our servers need to process the requests. That means the requests associated with the attack need to pass across the smaller, 1Gbps port on the server. From time to time, we've found that these attacks reached a large enough scale to overwhelm a 1Gbps port on one of our servers, making it a potential bottleneck.</p>
<p>Beyond raw bandwidth, the other bottleneck with some attacks centers on network interrupts. In most operating systems, every time a packet is received by a server, the network card generates an interrupt (known as an IRQ). An IRQ is effectively an instruction to the CPU to stop whatever else it's doing and deal with an event, in this case a packet arriving over the network. Each network adapter has multiple queues per port that receive these IRQs and then hands them to the server's CPU. The clock speed and driver efficiency in the network adapters, and<br/>message passing rate of the bus, effectively sets the maximum number of interrupts per second, and therefore packets per second, a server's network interface can handle.</p>
<p>In certain attacks, like large SYN floods which send a very high volume of very small packets, there can be plenty of bandwidth on a port but a CPU can be bottlenecked on IRQ handling. When this happens it can shut down a particular core on a CPU or, in the worst case if IRQs aren't properly balanced, shut down the whole CPU. To better deal with these attacks, we needed to find a way to more intelligently spread IRQs across more interfaces and, in turn, more CPU cores.</p>
<p>Both these problems are annoying if it affects the customer under attack, but it is unacceptable it spills over and affects customers who are not under attack. To ensure that would never happen, we needed to find a way to both increase network capacity and ensure that customer attacks were isolated from one another. To accomplish this we launched what we affectionately refer to in the office as "Project Bondage."</p>
<p><strong>Getting Into Bondage</strong></p>
<p>To deal with these challenges we started by implementing what is known as port bonding. The idea of port bonding is simple: use the resources of multiple ports in aggregate in order to support more traffic than any one port can on its own. We use a custom operating system based on the Debian line of Linux. Like most Linux varieties, our OS supports seven different port bonding modes:</p>
<ul>
<li>[0] Round-robin: Packets are transmitted sequentially through list of connections</li>
<li>[1] Active-backup: Only one connection is active, when it fails another is activated</li>
<li>[2] Balance-xor: This will ensure packets to a given destination from a given source will be the same over multiple connections</li>
<li>[3] Broadcast: Transmits everything over every active connection</li>
<li>[4] 802.3ad Dynamic Link Aggregation: Creates aggregation groups that share the same speed and duplex settings. Switches upstream must support 802.3ad.</li>
<li>[5] Balance-tlb: Adaptive transmit load balancing &mdash; outgoing traffic is balanced based on total amount being transmitted</li>
<li>[6] Balance-alb: Adaptive load balancing &mdash; includes balance-tlb and balances incoming traffic by using ARP negotiation to dynamically change the source MAC addresses of outgoing packets</li>
</ul>
<p>We use mode 4, 802.3ad Dynamic Link Aggregation. This requires switches that support 802.3ad (our workhorse switch is a Juniper&nbsp;4200, which does). Our switches are configured to send packets from each stream to the same network interface. If you want to experiment with port bonding yourself, the next section covers the technical details of exactly how we set it up.&nbsp;</p>
<p><strong>The Nitty Gritty</strong></p>
<p>Port bonding is configured on each server. It requires two Linux components that you can apt-get (assuming you're using a Debian-based Linux) if they're not already installed: ifenslave and ethtool. To initialize the bonding driver we use the following command:</p>
<div class="CodeRay">
<div class="code"><pre>modprobe bonding mode=4 miimon=100 xmit_hash_policy=1 lacp_rate=1</pre></div>
</div>
<p>Here's how that command breaks down:</p>
<ul>
<li><strong>mode=4</strong>: 802.3ad Dynamic Link Aggregation mode described above</li>
<li><strong>miimon=100</strong>: indicates that the devices are polled every 100ms to check for * connection changes, such as a link being down or a link duplex having changed.&nbsp;</li>
<li><strong>xmit_hash_policy=1</strong>: instructs the driver to spread the load over interfaces based on the source and destination IP address instead of MAC address</li>
<li><strong>lacp_rate=1</strong>: sets the rate for transmitting LACPDU packets, 0 is once every 30 seconds, 1 is every 1 second, which allows our network devices to automatically configure a single logical connection at the switch quickly</li>
</ul>
<p>After the bonding driver is initialized, we bring down the servers' network interfaces:</p>
<div class="CodeRay">
<div class="code"><pre>ifconfig eth0 down
ifconfig eth1 down</pre></div>
</div>
<p>We then bring up the bonding interface:</p>
<div class="CodeRay">
<div class="code"><pre>ifconfig bond0 192.168.0.2/24 up</pre></div>
</div>
<p>We then enslave (seriously, that's the term) the interfaces in the bond:</p>
<div class="CodeRay">
<div class="code"><pre>ifenslave bond0 eth0
ifenslave bond0 eth1</pre></div>
</div>
<p>Finally, we check the status of the bonded interface:</p>
<div class="CodeRay">
<div class="code"><pre>cat /proc/net/bonding/bond0</pre></div>
</div>
<p>From an application perspective, bonded ports appear as a single logical network interface with a higher maximum throughput. Since our switch recognizes and supports 802.3ad Dynamic Link Aggregation, we don't have to make any changes to its configuration in order for port bonding to work. In our case, we aggregate three ports (3Gbps) for handling external traffic and the remaining three ports (3Gbps) for handling intra-server traffic across our LAN.</p>
<p><strong>Working Out the Kinks</strong></p>
<p>Expanding the maximum effective capacity of each logical interface is half the battle. The other half is ensuring that network interrupts (IRQs) don't become a bottleneck. By default most Linux distributions rely on a service called irqbalance to set the CPU affinity of each IRQ queue. Unfortunately, we found that irqbalance does not effectively isolate each queue from overwhelming another on the same CPU. The problem with this is, because of the traffic we need to send from machine to machine, external attack traffic risked disrupting internal LAN traffic and affecting site performance beyond the customer under attack.</p>
<p>To solve this, the first thing we did was disable irqbalance:</p>
<div class="CodeRay">
<div class="code"><pre>/etc/init.d/irqbalance stop
update-rc.d irqbalance remove</pre></div>
</div>
<p>Instead, we explicitly setup IRQ handling to isolate our external and internal (LAN) networks. Each of our servers has two physical CPUs (G3 hardware uses a low-watt version of Intel Westmere line of CPUs) with six physical cores each. We use Intel's hyperthreading technology which effectively doubles the number of logical CPU cores: 12 per CPU or 24 per server.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Intel_x5645e" height="250" src="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-04-07/ksFmlqejekpAEftitkfGIqHhqDorhCyuzvrFgqHjqCIpflrHkhzojrtEGFbH/intel_x5645e.jpg.scaled500.jpg" width="250"/>
</div>
</p>
<p>Each port on our NICs has a number of queues to handle incoming requests. These are known as RSS (Receive Side Scaling) queues. Each port has 8 RSS queues, we have 6 1Gbps NIC ports per server, so a total of 48 RSS queues. These 48 RSS queues are allocated to the 24 cores, with 2 RSS queues per core. We divide the RSS queues between internal (LAN) traffic and external traffic and bind each type of traffic to one of the two server CPUs. This ensures that even large SYN floods that may affect a machine's ability to handle more external requests won't keep it from handling requests from other servers in the data center.</p>
<p><strong>The Results</strong></p>
<p>The net effect of these changes allows us to handle 30% larger SYN floods per server and increases our maximum throughput per site per server by 300%. Equally importantly, by custom tuning our IRQ handling, it has allowed us to ensure that customers under attack are isolated from those who are not while still delivering the maximum performance by fully utilizing all the gear in our network.</p>
<p>Near the end of 2012, our ops and networking teams sat down to spec our next generation of gear, incorporating everything we've learned over the previous year. One of the biggest changes we're making with G4 is the jump from 1Gbps network interfaces up to 10Gbps network interfaces on our switches and servers. Even without bonding, our tests of the new G4 gear show that it significantly increases both maximum throughput and IRQ handling. Or, put more succinctly: this next generation of servers is smokin' fast.</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-04-07/qkDBtfABGkyCAkImbuiipEjdtvuEdgwazggcCvHofegiqnmryCbemzeBqcaB/next-generation.jpg.scaled1000.jpg"><img alt="Next-generation" height="335" src="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-04-07/qkDBtfABGkyCAkImbuiipEjdtvuEdgwazggcCvHofegiqnmryCbemzeBqcaB/next-generation.jpg.scaled500.jpg" width="500" /></a>
</div>
</p>
<p>The first installations of the G4 gear is now in testing in a handful of our facilities. After testing, we plan to roll out worldwide over the coming months. We're already planning a detailed tour of the gear we chose, an explanation of the decisions we made, and performance benchmarks to show you how this next generation of gear is going to make CloudFlare's network even faster, safer, and smarter. That's a blog post I'm looking forward to writing. Stay tuned!</p>
	
</p>

<p><a href="http://blog.cloudflare.com/how-the-cloudflare-team-got-into-bondage-its">Permalink</a> 

	| <a href="http://blog.cloudflare.com/how-the-cloudflare-team-got-into-bondage-its#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="370" width="550" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-04-07/DhwAldFqkAruHepvqiHehettJJCoyJulEdipGnorkhxuyFgdirahmoqnwqmE/cat5-o-nine-tails.png">
        <media:thumbnail height="336" width="500" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-04-07/DhwAldFqkAruHepvqiHehettJJCoyJulEdipGnorkhxuyFgdirahmoqnwqmE/cat5-o-nine-tails.png.scaled500.png"/>
      </media:content>
      <media:content type="image/jpeg" height="640" width="480" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-04-07/rifaGwFyzhhpGwlvDBesrnwayGartdkIyGpillzzuiakfpDidCHnyvakgADy/cloudflare_servers.jpg">
        <media:thumbnail height="640" width="480" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-04-07/rifaGwFyzhhpGwlvDBesrnwayGartdkIyGpillzzuiakfpDidCHnyvakgADy/cloudflare_servers.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/jpeg" height="1524" width="2275" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-04-07/qkDBtfABGkyCAkImbuiipEjdtvuEdgwazggcCvHofegiqnmryCbemzeBqcaB/next-generation.jpg">
        <media:thumbnail height="335" width="500" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-04-07/qkDBtfABGkyCAkImbuiipEjdtvuEdgwazggcCvHofegiqnmryCbemzeBqcaB/next-generation.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/jpeg" height="250" width="250" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-04-07/ksFmlqejekpAEftitkfGIqHhqDorhCyuzvrFgqHjqCIpflrHkhzojrtEGFbH/intel_x5645e.jpg">
        <media:thumbnail height="250" width="250" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-04-07/ksFmlqejekpAEftitkfGIqHhqDorhCyuzvrFgqHjqCIpflrHkhzojrtEGFbH/intel_x5645e.jpg.scaled500.jpg"/>
      </media:content>
    </item>
    <item>
      <pubDate>Wed, 27 Mar 2013 09:35:00 -0700</pubDate>
      <title>The DDoS That Almost Broke the Internet</title>
      <link>http://blog.cloudflare.com/the-ddos-that-almost-broke-the-internet</link>
      <guid>http://blog.cloudflare.com/the-ddos-that-almost-broke-the-internet</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<img alt="Massive_attack" height="377" src="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-27/FkpGIvzyFhFrpfErjufvAypfDuCDstbzkzaEgAiEsdyGBozsodJinyEtynie/massive_attack.jpg.scaled500.jpg" width="500" />
</div>
</p>
<p>The <em>New York Times</em>&nbsp;this morning published a story about the <a href="http://www.nytimes.com/2013/03/27/technology/internet/online-dispute-becomes-internet-snarling-attack.html?" target="_blank">Spamhaus DDoS attack and how CloudFlare helped mitigate it and keep the site online</a>. The <em>Times</em> calls the attack the largest known DDoS attack ever on the Internet. We <a href="http://blog.cloudflare.com/the-ddos-that-knocked-spamhaus-offline-and-ho" target="_blank">wrote about the attack last week</a>. At the time, it was a large attack, sending 85Gbps of traffic. Since then, the attack got much worse. Here are some of the technical details of what we've seen.</p>
<p><strong>Growth Spurt</strong></p>
<p>On Monday, March 18, 2013 Spamhaus contacted CloudFlare regarding an attack they were seeing against their website <a href="http://www.spamhaus.org" target="_blank">spamhaus.org</a>. They signed up for CloudFlare and we quickly mitigated the attack. The attack, initially, was approximately 10Gbps generated largely from open DNS recursors. On March 19, the attack increased in size, peaking at approximately 90Gbps. The attack fluctuated between 90Gbps and 30Gbps until 01:15 UTC on on March 21.</p>
<p>The attackers were quiet for a day. Then, on March 22 at 18:00 UTC, the attack resumed, peaking at 120Gbps of traffic hitting our network. As we discussed in the previous blog post, CloudFlare uses Anycast technology which spreads the load of a distributed attack across all our data centers. This allowed us to mitigate the attack without it affecting Spamhaus or any of our other customers. The attackers ceased their attack against the Spamhaus website four hours after it started.</p>
<p>Other than the scale, which was already among the largest DDoS attacks we've seen, there was nothing particularly unusual about the attack to this point. Then the attackers changed their tactics. Rather than attacking our customers directly, they started going after the network providers CloudFlare uses for bandwidth. More on that in a second, first a bit about how the Internet works.</p>
<p><strong>Peering on the Internet</strong></p>
<p>The "inter" in Internet refers to the fact that it is a collection of independent networks connected together. CloudFlare runs a network, Google runs a network, and bandwidth providers like Level3, AT&amp;T, and Cogent run networks. These networks then interconnect through what are known as peering relationships.</p>
<p>When you surf the web, your browser sends and receives packets of information. These packets are sent from one network to another. You can see this by running a traceroute. Here's one from <a href="http://www.slac.stanford.edu/cgi-bin/nph-traceroute.pl" target="_blank">Stanford University's network</a> to the New York Times' website (nytimes.com):</p>
<div class="CodeRay">
  <div class="code"><pre>1  rtr-servcore1-serv01-webserv.slac.stanford.edu (134.79.197.130)  0.572 ms
 2  rtr-core1-p2p-servcore1.slac.stanford.edu (134.79.252.166)  0.796 ms
 3  rtr-border1-p2p-core1.slac.stanford.edu (134.79.252.133)  0.536 ms
 4  slac-mr2-p2p-rtr-border1.slac.stanford.edu (192.68.191.245)  25.636 ms
 5  sunncr5-ip-a-slacmr2.es.net (134.55.36.21)  3.306 ms
 6  eqxsjrt1-te-sunncr5.es.net (134.55.38.146)  1.384 ms
 7  xe-0-3-0.cr1.sjc2.us.above.net (64.125.24.1)  2.722 ms
 8  xe-0-1-0.mpr1.sea1.us.above.net (64.125.31.17)  20.812 ms
 9  209.249.122.125 (209.249.122.125)  21.385 ms</pre></div>
</div>

<p>There are three networks in the above traceroute: stanford.edu, es.net, and above.net. The request starts at Stanford. Between lines 4 and 5 it passes from Stanford's network to their peer es.net. Then, between lines 6 and 7, it passes from es.net to above.net, which appears to provide hosting for the New York Times. This means Stanford has a peering relationship with ES.net. ES.net has a peering relationship with Above.net. And Above.net provides connectivity for the New York Times.</p>
<p>CloudFlare connects to a large number of networks. You can get a sense of some, although not all, of the networks we peer with through a tool like <a href="http://bgp.he.net/AS13335#_peers" target="_blank">Hurricane Electric's BGP looking glass</a>. CloudFlare connects to peers in two ways. First, we connect directly to certain large carriers and other networks to which we send a large amount of traffic. In this case, we connect our router directly to the router at the border of the other network, usually with a piece of fiber optic cable. Second, we connect to what are known as Internet Exchanges, IXs for short, where a number of networks meet in a central point.</p>
<p>Most major cities have an IX. The model for IXs are different in different parts of the world. Europe runs some of the most robust IXs, and CloudFlare connects to several of them including LINX (the London Internet Exchange), AMS-IX (the Amsterdam Internet Exchange), and DE-CIX (the Frankfurt Internet Exchange), among others. The major networks that make up the Internet --Google, Facebook Yahoo, etc. -- connect to these same exchanges to pass traffic between each other efficiently. When the Spamhaus attacker realized he couldn't go after CloudFlare directly, he began targeting our upstream peers and exchanges.</p>
<p><strong>Headwaters</strong></p>
<p>Once the attackers realized they couldn't knock CloudFlare itself offline even with more than 100Gbps of DDoS traffic, they went after our direct peers. In this case, they attacked the providers from whom CloudFlare buys bandwidth. We, primarily, contract with what are known as Tier 2 providers for CloudFlare's paid bandwidth. These companies peer with other providers and also buy bandwidth from so-called Tier 1 providers.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Peer_pressure" height="472" src="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-03-27/clqiHAuohBfofwffxtDoEHjnHnwvFBxJgbJahEpHbnomvoIbxCBukDwbEFEs/peer_pressure.png.scaled500.png" width="500" />
</div>
</p>
<p>There are <a href="http://en.wikipedia.org/wiki/Tier_1_network" target="_blank">approximately a dozen Tier 1 providers</a> on the Internet. The nature of these providers is that they don't buy bandwidth from anyone. Instead, they engage in what is known as settlement-free peering with the other Tier 1 providers. Tier 2 providers interconnect with each other and then buy bandwidth from the Tier 1 providers in order to ensure they can connect to every other point on the Internet. At the core of the Internet, if all else fails, it is these Tier 1 providers that ensure that every network is connected to every other network. If one of them fails, it's a big deal.</p>
<p>Anycast means that if the attacker attacked the last step in the traceroute then their attack would be spread across CloudFlare's worldwide network, so instead they attacked the second to last step which concentrated the attack on one single point. This wouldn't cause a network-wide outage, but it could potentially cause regional problems.</p>
<p>We carefully select our bandwidth providers to ensure they have the ability to deal with attacks like this. Our direct peers quickly filtered attack traffic at their edge. This pushed the attack upstream to their direct peers, largely Tier 1 networks. Tier 1 networks don't buy bandwidth from anyone, so the majority of the weight of the attack ended up being carried by them. While we don't have direct visibility into the traffic loads they saw, we have been told by one major Tier 1 provider that they saw more than 300Gbps of attack traffic related to this attack. That would make this attack one of the largest ever reported.</p>
<p>The challenge with attacks at this scale is they risk overwhelming the systems that link together the Internet itself. The largest routers that you can buy have, at most, 100Gbps ports. It is possible to bond more than one of these ports together to create capacity that is greater than 100Gbps however, at some point, there are limits to how much these routers can handle. If that limit is exceeded then the network becomes congested and slows down.</p>
<p>Over the last few days, as these attacks have increased, we've seen congestion across several major Tier 1s, primarily in Europe where most of the attacks were concentrated, that would have affected hundreds of millions of people even as they surfed sites unrelated to Spamhaus or CloudFlare. If the Internet felt a bit more sluggish for you over the last few days in Europe, this may be part of the reason why.</p>
<p><strong>Attacks on the IXs</strong></p>
<p>In addition to CloudFlare's direct peers, we also connect with other networks over the so-called Internet Exchanges (IXs). These IXs are, at their most basic level, switches into which multiple networks connect and can then pass bandwidth. In Europe, these IXs are run as non-profit entities and are considered critical infrastructure. They interconnect hundreds of the world's largest networks including CloudFlare, Google, Facebook, and just about every other major Internet company.</p>
<p>Beyond attacking CloudFlare's direct peers, the attackers also attacked the core IX infrastructure on the London Internet Exchange (LINX), the Amsterdam Internet Exchange (AMS-IX), the Frankfurt Internet Exchange (DE-CIX), and the Hong Kong Internet Exchange (HKIX). <a name="linxedit"></a>From our perspective, the attacks had the largest effect on LINX which caused impact over the exchange and LINX's systems that monitor the exchange, as visible through the drop in traffic recorded by their monitoring systems. (Corrected: see below for original phrasing.)</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-03-27/CpsvlGEBGEEGtchddisFijJdaJCwnICtnixuDDcJEGhcnspICfzpEpgabCGG/linx_traffic.png.scaled1000.png"><img alt="Linx_traffic" height="206" src="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-27/CpsvlGEBGEEGtchddisFijJdaJCwnICtnixuDDcJEGhcnspICfzpEpgabCGG/linx_traffic.png.scaled500.png" width="500" /></a>
</div>
</p>
<p>The congestion impacted many of the networks on the IXs, including CloudFlare's. As problems were detected on the IX, we would route traffic around them. However, several London-based CloudFlare users reported intermittent issues over the last several days. This is the root cause of those problems.</p>
<p>The attacks also exposed some vulnerabilities in the architecture of some IXs. We, along with many other network security experts, worked with the team at LINX to better secure themselves. In doing so, we developed a list of best practices for any IX in order to make them less vulnerable to attacks.</p>
<p>Two specific suggestions to limit attacks like this involve making it more difficult to attack the IP addresses that members of the IX use to interchange traffic between each other. We are working with IXs to ensure that: 1) these IP addresses should not be announced as routable across the public Internet; and 2) p<span style="font-family: arial, sans-serif;">ackets destined to these IP addresses should only be&nbsp;</span><span style="font-family: arial, sans-serif;">permitted from other IX IP addresses. We've been very impressed with the team at LINX and how quickly they've worked to implement these changes and add additional security to their IX and are hopeful other IXs will quickly follow their lead.</span></p>
<p><strong><span style="font-family: arial, sans-serif;">The Full Impact of the Open Recursor Problem</span></strong></p>
<p><span style="font-family: arial, sans-serif;">At the bottom of this attack we once again find the problem of open DNS recursors. The attackers were able to generate more than 300Gbps of traffic likely with a network of their own that only had access 1/100th of that amount of traffic themselves. We've written about how these mis-configured DNS recursors as a<a href="http://blog.cloudflare.com/deep-inside-a-dns-amplification-ddos-attack" target="_blank"> bomb waiting to go off</a> that literally threatens the stability of the Internet itself. We've now seen an attack that begins to illustrate the full extent of the problem.</span></p>
<p><span style="font-family: arial, sans-serif;">While lists of open recursors have been passed around on network security lists for the last few years, on Monday the full extent of the problem was, for the first time, made public. The <a href="http://openresolverproject.org" target="_blank">Open Resolver Project</a> made available the full list of the 21.7 million open resolvers online in an effort to shut them down.</span></p>
<p><span style="font-family: arial, sans-serif;">We'd debated doing the same thing ourselves for some time but worried about the collateral damage of what would happen if such a list fell into the hands of the bad guys. The last five days have made clear that the bad guys have the list of open resolvers and they are getting increasingly brazen in the attacks they are willing to launch. We are in full support of the Open Resolver Project and believe it is incumbent on all network providers to work with their customers to close any open resolvers running on their networks.</span></p>
<p><span style="font-family: arial, sans-serif;"><div class='p_embed p_image_embed'>
<a href="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-27/vbpodwjusyHFmBnECfbbbbFdkyeDgqqlhzCdwxqdwmBidimajjwuJibJECnC/bazookas.jpg.scaled1000.jpg"><img alt="Bazookas" height="289" src="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-03-27/vbpodwjusyHFmBnECfbbbbFdkyeDgqqlhzCdwxqdwmBidimajjwuJibJECnC/bazookas.jpg.scaled500.jpg" width="500" /></a>
</div>
</span></p>
<p><span style="font-family: arial, sans-serif;">Unlike traditional botnets which could only generate limited traffic because of the modest Internet connections and home PCs they typically run on, these open resolvers are typically running on big servers with fat pipes. They are like bazookas and the events of the last week have shown the damage they can cause. What's troubling is that, compared with what is possible, this attack may prove to be relatively modest.</span></p>
<p><span style="font-family: arial, sans-serif;">As someone in charge of DDoS mitigation at one of the Internet giants emailed me this weekend: "I've often said we don't have to prepare for the largest-possible attack, we just have to prepare for the largest attack the Internet can send without causing massive collateral damage to others. It looks like you've reached that point, so... congratulations!"</span></p>
<p><span style="font-family: arial, sans-serif;">At CloudFlare one of our goals is to make DDoS something you only read about in the history books. We're proud of how our network held up under such a massive attack and are working with our peers and partners to ensure that the Internet overall can stand up to the threats it faces.</span></p>
<p><strong>Correction</strong>: The original sentence about the impact on LINX was "From our perspective, the attacks had the largest effect on LINX which for a little over an hour on March 23 saw the infrastructure serving more than half of the usual 1.5Tbps of peak traffic fail." That was not well phrased, and has been edited, with notation in place.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/the-ddos-that-almost-broke-the-internet">Permalink</a> 

	| <a href="http://blog.cloudflare.com/the-ddos-that-almost-broke-the-internet#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/jpeg" height="377" width="500" url="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-03-27/FkpGIvzyFhFrpfErjufvAypfDuCDstbzkzaEgAiEsdyGBozsodJinyEtynie/massive_attack.jpg">
        <media:thumbnail height="377" width="500" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-27/FkpGIvzyFhFrpfErjufvAypfDuCDstbzkzaEgAiEsdyGBozsodJinyEtynie/massive_attack.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/png" height="472" width="500" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-03-27/clqiHAuohBfofwffxtDoEHjnHnwvFBxJgbJahEpHbnomvoIbxCBukDwbEFEs/peer_pressure.png">
        <media:thumbnail height="472" width="500" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-03-27/clqiHAuohBfofwffxtDoEHjnHnwvFBxJgbJahEpHbnomvoIbxCBukDwbEFEs/peer_pressure.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="228" width="553" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-27/CpsvlGEBGEEGtchddisFijJdaJCwnICtnixuDDcJEGhcnspICfzpEpgabCGG/linx_traffic.png">
        <media:thumbnail height="206" width="500" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-27/CpsvlGEBGEEGtchddisFijJdaJCwnICtnixuDDcJEGhcnspICfzpEpgabCGG/linx_traffic.png.scaled500.png"/>
      </media:content>
      <media:content type="image/jpeg" height="291" width="504" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-03-27/vbpodwjusyHFmBnECfbbbbFdkyeDgqqlhzCdwxqdwmBidimajjwuJibJECnC/bazookas.jpg">
        <media:thumbnail height="289" width="500" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-03-27/vbpodwjusyHFmBnECfbbbbFdkyeDgqqlhzCdwxqdwmBidimajjwuJibJECnC/bazookas.jpg.scaled500.jpg"/>
      </media:content>
    </item>
    <item>
      <pubDate>Fri, 22 Mar 2013 14:53:00 -0700</pubDate>
      <title>Page Rules Reordering Now Available</title>
      <link>http://blog.cloudflare.com/page-rules-reordering-now-available</link>
      <guid>http://blog.cloudflare.com/page-rules-reordering-now-available</guid>
      <description>
        <![CDATA[<p>
	<p><span style="color: #000000; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; line-height: 18.1875px; font-size: small;">Page Rules are powerful tools for controlling how CloudFlare works on your site on a page-by-page basis. Customers customize CloudFlare with</span><span style="font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: small; line-height: 18.1875px;">&nbsp;Page Rules </span><span style="font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: small; line-height: 18.1875px;">based on their specific needs, including changing or extending caching, forwarding URLs, or disabling certain features for specific pages or directories.</span></p>
<p><span style="font-size: small; color: #000000;">Today, we're making managing Page Rules even easier with Page Rules reordering.</span></p>
<p><span style="font-size: small; color: #000000;"> Page Rules are applied in the order they appear in your CloudFlare dashboard, from top to bottom. The order matters, since the first Page Rule to match the request is applied. So, if you want to apply aggressive caching to a specific set of pages but exclude caching on one login or admin URL, you'd put the exclude caching Page Rule above the aggressive caching Page Rule.</span></p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-03-22/qbvgzdpzjnvEvxeApacbBuAGvavdxjkqHdBjkccgplqvhxjDIHmdcajnAazj/Page_Rules_screenshot.tiff.scaled1000.jpg"><img alt="Page_rules_screenshot" height="221" src="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-22/qbvgzdpzjnvEvxeApacbBuAGvavdxjkqHdBjkccgplqvhxjDIHmdcajnAazj/Page_Rules_screenshot.tiff.scaled500.jpg" width="500" /></a>
</div>
</p>
<p><span style="font-size: small; color: #000000;">Before today, reordering Page Rules was cumbersome. You had to delete and re-add Page Rules in the order that you wanted them to be applied. Now, you can easily reorder Page Rules by clicking on the icon to the left and dragging them up or down.&nbsp;</span></p>
<p>If you're not familiar with Page Rules or not sure how to use them, review this tutorial:</p>
<p><a href="https://support.cloudflare.com/entries/22576178-Is-there-a-tutorial-for-PageRules-" title="https://support.cloudflare.com/entries/22576178-Is-there-a-tutorial-for-PageRules-" target="_blank">https://support.cloudflare.com/entries/22576178-Is-there-a-tutorial-for-PageRules-</a></p>
<p>Page Rules are found in the menu available under the gear icon in My Websites, as shown in this screenshot.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Page-rules-in-menu" height="313" src="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-22/CFigjfnyJofEJGmCDbmIwyFbDzvstCCbFwbyppJdkzziaIppgvmjarCjqeJm/page-rules-in-menu.png.scaled500.png" width="355" />
</div>
</p>
<p>Reordering feature is available <strong>now</strong> for all customers, with the number of Page Rules set by plan type<span style="font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: small; line-height: 18.1875px;">: Free domains get 3, Pro domains get 20, Business domains get 50, and Enterprise domains get a custom number.</span></p>
<p>&nbsp;</p>
	
</p>

<p><a href="http://blog.cloudflare.com/page-rules-reordering-now-available">Permalink</a> 

	| <a href="http://blog.cloudflare.com/page-rules-reordering-now-available#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725124/IMG_0043.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbv0Udt1n</posterous:profileUrl>
        <posterous:firstName>Michelle</posterous:firstName>
        <posterous:lastName>Zatlyn</posterous:lastName>
        <posterous:nickName>Michelle Z.</posterous:nickName>
        <posterous:displayName>Michelle Zatlyn</posterous:displayName>
      </posterous:author>
      <media:content type="image/tiff" height="445" width="1007" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-03-22/qbvgzdpzjnvEvxeApacbBuAGvavdxjkqHdBjkccgplqvhxjDIHmdcajnAazj/Page_Rules_screenshot.tiff">
        <media:thumbnail height="221" width="500" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-22/qbvgzdpzjnvEvxeApacbBuAGvavdxjkqHdBjkccgplqvhxjDIHmdcajnAazj/Page_Rules_screenshot.tiff.scaled500.jpg"/>
      </media:content>
      <media:content type="image/png" height="313" width="355" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-03-22/CFigjfnyJofEJGmCDbmIwyFbDzvstCCbFwbyppJdkzziaIppgvmjarCjqeJm/page-rules-in-menu.png">
        <media:thumbnail height="313" width="355" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-22/CFigjfnyJofEJGmCDbmIwyFbDzvstCCbFwbyppJdkzziaIppgvmjarCjqeJm/page-rules-in-menu.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Thu, 21 Mar 2013 18:26:39 -0700</pubDate>
      <title>Leading Experts Weigh in on Industry Trends at Parallels Summit 2013</title>
      <link>http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a</link>
      <guid>http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a</guid>
      <description>
        <![CDATA[<p>
	<p>Last month we attended Parallels Summit to meet with new and existing partners, and hear from leading experts on the latest industry trends. We had a great time giving limo rides to summit attendees and hosting &ldquo;Conversations with CloudFlare&rdquo; at our booth.</p>
<p>For these conversations, <a href="http://www.cloudflare.com/">CloudFlare</a> co-founder<a href="https://www.twitter.com/zatlyn"> Michelle Zatlyn</a> sat down with 15 leading experts in the hosting and service provider industry. Their conversations were captured live and offer insight into the latest trends and news in hosting and web services.</p>
<p>Below are highlights and video clips from five of those conversations. We will be posting the rest of the videos next week.</p>
<p><br /><strong><a href="https://www.twitter.com/ParallelsCloud" target="_blank">Dan Havens</a>, VP of Sales at Parallels<br /></strong><em>When it comes to the adoption of Plesk 11, Dan says it&rsquo;s all about the gateway. &nbsp;Listen in as he discusses Plesk 11 benefits, changes to the hosting space in the last few years, and what country Parallels is really big in right now...</em></p>
<p><iframe src="http://player.vimeo.com/video/61677896?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p><strong><br /><a href="https://twitter.com/srenkema" target="_blank">Sam Renkema</a>, CEO of Spam Experts</strong><br /><em>Sam knows spam. As the CEO of one of the most widely deployed packages offered in the APS catalog, Sam has seen all the trends in email spam. Sam tells us what hosting providers are most concerned about and what the newest sector of email spam looks like...</em></p>
<p><iframe src="http://player.vimeo.com/video/61745929?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p><strong><br /><a href="https://twitter.com/mhouwen" target="_blank">Marco Howen</a>, CEO of LuxCloud<br /></strong><em>&ldquo;We have to adapt. If we&rsquo;re not doing it, we won&rsquo;t sell&rdquo; - Marco emphasizes ease of adoption and making UIs as simple as possible. Oh, and LuxCloud is a globally impacting company with an employee headcount that might surprise you...</em></p>
<p><iframe src="http://player.vimeo.com/video/61677899?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p>&nbsp;</p>
<p><strong>Terry Gerstner, VP of Marketing at MobeeArt<br /></strong><em>MobeeArt is like mobile magic, for your website. MobeeArt automatically converts websites to mobile. Sound too good to be true? It&rsquo;s not. And they are announcing a feature that&rsquo;s going to make running a mobile website even better...</em></p>
<p><iframe src="http://player.vimeo.com/video/61316381?portrait=0" frameborder="0" height="283" width="500"></iframe></p>
<p><br /><strong><br /><a href="https://www.twitter.com/blacknight	" target="_blank">Michele Nelyon</a>, CEO of Blacknight</strong><br /><em>Michele has his finger on the pulse of Internet policy. There are big issues the hosting industry will be facing over the next few years. Michele discusses these issues and what we can do to help...</em></p>
<p><a href="https://vimeo.com/61743347"><iframe src="http://player.vimeo.com/video/61743347?portrait=0" frameborder="0" height="283" width="500"></iframe></a></p>
<p><strong style="font-family: Times; font-size: medium; font-weight: normal;"><br /></strong></p>
<p>&nbsp;</p>
	
</p>

<p><a href="http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a">Permalink</a> 

	| <a href="http://blog.cloudflare.com/leading-experts-weigh-in-on-industry-trends-a#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/1122929/CloudFlare4-16.jpeg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/he6rIlp3ceWH8</posterous:profileUrl>
        <posterous:firstName>Kristin</posterous:firstName>
        <posterous:lastName>Tarr</posterous:lastName>
        <posterous:nickName>ktarr</posterous:nickName>
        <posterous:displayName>Kristin Tarr</posterous:displayName>
      </posterous:author>
    </item>
    <item>
      <pubDate>Wed, 20 Mar 2013 11:26:00 -0700</pubDate>
      <title>The DDoS That Knocked Spamhaus Offline (And How We Mitigated It)</title>
      <link>http://blog.cloudflare.com/the-ddos-that-knocked-spamhaus-offline-and-ho</link>
      <guid>http://blog.cloudflare.com/the-ddos-that-knocked-spamhaus-offline-and-ho</guid>
      <description>
        <![CDATA[<p>
	<p>At CloudFlare, we deal with large DDoS attacks every day. Usually, these attacks are directed at large companies or organizations that are reluctant to talk about their details. It's fun, therefore, whenever we have a customer that is willing to let us tell the story of an attack they saw and how we mitigated it. This is one of those stories.</p>
<p><strong>Spamhaus</strong></p>
<p>Yesterday, Tuesday, March 19, 2013, CloudFlare was contacted by the non-profit anti-spam organization <a href="http://www.spamhaus.org/" target="_blank">Spamhaus</a>. They were suffering a large DDoS attack against their website and asked if we could help mitigate the attack.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Spamhaus_logo" height="159" src="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-20/oqehwirAwbrxpqdJBJenzxdmAefduBHaGCsdwBGjbEEEAlzqablaIuECumpB/spamhaus_logo.jpg.scaled500.jpg" width="305" />
</div>
</p>
<p>Spamhaus provides one of the key backbones that underpins much of the anti-spam filtering online. Run by a tireless team of volunteers, Spamhaus patrols the Internet for spammers and publishes a list of the servers they use to send their messages in order to empower email system administrators to filter unwanted messages. Spamhaus's services are so pervasive and important to the operation of the Internet's email architecture that, when a <a href="http://www.theregister.co.uk/2011/09/05/spamhaus_e360_insight_lawsuit/" target="_blank">lawsuit threatened to shut the service down</a>, industry experts testified [<a href="http://app.quickblogcast.com/files/31236-29497/spamhaus_amicus.pdf" target="_blank">PDF</a>, full disclosure: I wrote the brief back in the day] that doing so risked literally breaking email since Spamhaus is directly or indirectly responsible for filtering as much as 80% of daily spam messages.</p>
<p>Beginning on March 18, the Spamhaus site <a href="https://isc.sans.edu/diary/Spamhaus+DDOS/15427" target="_blank">came under attack</a>. The attack was large enough that the Spamhaus team wasn't sure of its size when they contacted us. It was sufficiently large to fully saturate their connection to the rest of the Internet and knock their site offline. These very large attacks, which are known as Layer 3 attacks, are difficult to stop with any on-premise solution. Put simply: if you have a router with a 10Gbps port, and someone sends you 11Gbps of traffic, it doesn't matter what intelligent software you have to stop the attack because your network link is completely saturated.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Burst_pipe" height="259" src="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-20/IerauHyfaagukmpyqyfBqClrfybyjBEHmrBknaJfpGlmItuGbfakiylsCfxy/burst_pipe.jpg.scaled500.jpg" width="400" />
</div>
</p>
<p>While we don't know who was behind this attack, Spamhaus has made plenty of enemies over the years. Spammers aren't always the most lovable of individuals and Spamhaus has been threatened, sued, and DDoSed regularly. Spamhaus's blocklists are distributed via DNS and there is a long list of volunteer organizations that mirror their DNS infrastructure in order to ensure it is resilient to attacks. The website, however, was unreachable.</p>
<p><strong>Filling Up the Series of Tubes</strong></p>
<p>Very large Layer 3 attacks are nearly always originated from a number of sources. These many sources each send traffic to a single Internet location, effectively creating a tidal wave that overwhelms the target's resources. In this sense, the attack is distributed (the first D in DDoS -- Distributed Denial of Service). The sources of attack traffic can be a group of individuals working together (e.g., the Anonymous LOIC model, although this is Layer 7 traffic and even at high volumes usually much smaller in volume than other methods), a botnet of compromised PCs, a botnet of compromised servers, <a href="http://blog.cloudflare.com/deep-inside-a-dns-amplification-ddos-attack" target="_blank">misconfigured DNS resolvers</a>, or even <a href="http://internetcensus2012.bitbucket.org/paper.html" target="_blank">home Internet routers with weak passwords</a>.</p>
<p>Since an attacker attempting to launch a Layer 3 attack doesn't care about receiving a response to the requests they send, the packets that make up the attack do not have to be accurate or correctly formatted. Attackers will regularly spoof all the information in the attack packets, including the source IP, making it look like the attack is coming from a virtually infinite number of sources. Since packets data can be fully randomized, using techniques like IP filtering even upstream becomes virtually useless.</p>
<p>Spamhaus signed up for CloudFlare on Tuesday afternoon and we immediately mitigated the attack, making the site once again reachable. (More on how we did that below.) Once on our network, we also began recording data about the attack. At first, the attack was relatively modest (around 10Gbps). There was a brief spike around 16:30 UTC, likely a test, that lasted approximately 10 minutes. Then, around 21:30 UTC, the attackers let loose a very large wave.</p>
<p>The graph below is generated from bandwidth samples across a number of the routers that sit in front of servers we use for DDoS scrubbing. The green area represents in-bound requests and the blue line represents out-bound responses. While there is always some attack traffic on our network, it's easy to see when the attack against Spamhaus started and then began to taper off around 02:30 UTC on March 20, 2013. As I'm writing this at 16:15 UTC on March 20, 2013, it appears the attack is picking up again.</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-03-20/lwhhFndFymGjBljDbgFgkFHdHxtDanAdlqGlyqeHqAAmFJgkqIbEHsoIfsmc/spamhaus_ddos_attack.png.scaled1000.png"><img alt="Spamhaus_ddos_attack" height="159" src="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-20/lwhhFndFymGjBljDbgFgkFHdHxtDanAdlqGlyqeHqAAmFJgkqIbEHsoIfsmc/spamhaus_ddos_attack.png.scaled500.png" width="500" /></a>
</div>
</p>
<p><strong>How to Generate a 75Gbps DDoS</strong></p>
<p>The largest source of attack traffic against Spamhaus came from DNS reflection. I've <a href="http://blog.cloudflare.com/deep-inside-a-dns-amplification-ddos-attack" target="_blank">written about these attacks before</a> and in the last year they have become the source of the largest Layer 3 DDoS attacks we see (sometimes well exceeding 100Gbps). Open DNS resolvers are quickly becoming the scourge of the Internet and the size of these attacks will only continue to rise until all providers make a <a href="http://blog.cloudflare.com/good-news-open-dns-resolvers-are-getting-clos" target="_blank">concerted effort to close them</a>. (It also makes sense to implement&nbsp;<a href="http://tools.ietf.org/html/bcp38" target="_blank">BCP-38</a>, but that's a topic for another post another time.)</p>
<p>The basic technique of a DNS reflection attack is to send a request for a large DNS zone file with the source IP address spoofed to be the intended victim to a large number of open DNS resolvers. The resolvers then respond to the request, sending the large DNS zone answer to the intended victim. The attackers' requests themselves are only a fraction of the size of the responses, meaning the attacker can effectively amplify their attack to many times the size of the bandwidth resources they themselves control.&nbsp;</p>
<p>In the Spamhaus case, the attacker was sending requests for the DNS zone file for ripe.net to open DNS resolvers. The attacker spoofed the CloudFlare IPs we'd issued for Spamhaus as the source in their DNS requests. The open resolvers responded with DNS zone file, generating collectively approximately 75Gbps of attack traffic. The requests were likely approximately 36 bytes long (e.g. dig ANY ripe.net @X.X.X.X +edns=0 +bufsize=4096, where X.X.X.X is replaced with the IP address of an open DNS resolver) and the response was approximately 3,000 bytes, translating to a 100x amplification factor.</p>
<p>We recorded over 30,000 unique DNS resolvers involved in the attack. This translates to each open DNS resolver sending an average of 2.5Mbps, which is small enough to fly under the radar of most DNS resolvers. Because the attacker used a DNS amplification, the attacker only needed to control a botnet or cluster of servers to generate 750Mbps -- which is possible with a small sized botnet or a handful of AWS instances. It is worth repeating: open DNS resolvers are the scourge of the Internet and these attacks will become more common and large until service providers take serious efforts to close them.</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-20/rnFosdBxIJCjCexwHzDthcAawvbeqwtqFkyrImksJkqdgBrDwGAhGskhGadm/im_under_attack.jpg.scaled1000.jpg"><img alt="Im_under_attack" height="213" src="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-20/rnFosdBxIJCjCexwHzDthcAawvbeqwtqFkyrImksJkqdgBrDwGAhGskhGadm/im_under_attack.jpg.scaled500.jpg" width="500" /></a>
</div>
</p>
<p><strong>How You Mitigate a 75Gbps DDoS</strong></p>
<p>While large Layer 3 attacks are difficult for an on-premise DDoS solution to mitigate, CloudFlare's network was specifically designed from the beginning to stop these types of attacks. We make heavy use of Anycast. That means the same IP address is announced from every one of our 23 worldwide data centers. The network itself <a href="http://blog.cloudflare.com/cloudflares-architecture-eliminating-single-p" target="_blank">load balances requests</a> to the nearest facility. Under normal circumstances, this helps us ensure a visitor is routed to the nearest data center on our network.</p>
<p>When there's an attack, Anycast serves to effectively dilute it by spreading it across our facilities. Since every data center announces the same IP address for any CloudFlare customer, traffic cannot be concentrated in any one location. Instead of the attack being many-to-one, it becomes many-to-many with no single point on the network acting as a bottleneck.</p>
<p>Once diluted, the attack becomes relatively easy to stop at each of our data centers. Because CloudFlare acts as a virtual shield in front of our customers sites, with Layer 3 attacks none of the attack traffic reaches the customer's servers. Traffic to Spamhaus's network dropped to below the levels when the attack started as soon as they signed up for our service.</p>
<p><strong>Other Noise</strong></p>
<p>While the majority of the traffic involved in the attack was DNS reflection, the attacker threw in a few other attack methods as well. One was a so-called ACK reflection attack. When a TCP connection is established there is a handshake. The server initiating the TCP session first sends a SYN (for synchronize) request to the receiving server. The receiving server responds with an ACK (for acknowledge). After that handshake, data can be exchanged.</p>
<p>In an ACK reflection, the attacker sends a number of SYN packets to servers with a spoofed source IP address pointing to the intended victim. The servers then respond to the victim's IP with an ACK. Like the DNS reflection attack, this disguises the source of the attack, making it appear to come from legitimate servers. However, unlike the DNS reflection attack, there is no amplification factor: the bandwidth from the ACKs is symmetrical to the bandwidth the attacker has to generate the SYNs. CloudFlare is configured to drop unmatched ACKs, which mitigates these types of attacks.</p>
<p>Whenever we see one of these large attacks, network operators will write to us upset that we are attacking their infrastructure with abusive DNS queries or SYN floods. In fact, it is their infrastructure that is being used to reflect an attack at us. By working with and educating network operators, they clean up their network which helps to solve the root cause of these large attacks.</p>
<p><strong>History Repeats Itself</strong></p>
<p>Finally, it's worth noting how similar this battle against DDoS attacks and open DNS relays is with Spamhaus's original fight. If DDoS is the network scourge of tomorrow, spam was its clear predecessor. Paul Vixie, <a href="http://en.wikipedia.org/wiki/DNSBL" target="_blank">the father of the DNSBL</a>, set out in 1997 to use DNS to help shut down the spam source of the day: open email relays. These relays were being used to disguise the origin of spam messages, making them more difficult to block. What was needed was a list of mail relays that mail serves could query against and decide whether to accept messages.</p>
<p><div class='p_embed p_image_embed'>
<img alt="History_repeats_itself" height="131" src="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-03-20/snzgfjctdrklJyqvsunggCHefAIHfqbIpCHeuIjCcvibdgujIHDsbIfAkFeu/history_repeats_itself.png.scaled500.png" width="481" />
</div>
</p>
<p>While it wasn't originally designed with the idea in mind, DNS proved a highly scalable and efficient means to distribute a queryable list of open mail relays that email service providers could use to block unwanted messages. Spamhaus arose as one of the most respected and widely used DNSBLs, effectively blocking a huge percentage of daily spam volume.</p>
<p>As open mail relays were shut, spammers turned to virus writers to create botnets that could be used to relay spam. Spamhaus expanded their operations to list the IPs of known botnets, trying to stay ahead of spammers. CloudFlare's own history grew out of <a href="http://www.projecthoneypot.org/" target="_blank">Project Honey Pot</a>, which started as an automated service to track the resources used by spammers and publishes the HTTP:BL.</p>
<p>Today, as Spamhaus's success has eroded the business model of spammers, botnet operators are increasingly renting their networks to launch DDoS attacks. At the same time, DNSBLs proved that there were many functions that the DNS protocol could be used for, encouraging many people to tinker with installing their own DNS resolvers. Unfortunately, these DNS resolvers are often mis-configured and left open to abuse, making them the DDoS equivalent of the open mail relay.</p>
<p>If you're running a network, take a second to make sure you've closed any open resolvers before DDoS explodes into an even worse problem than it already is.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/the-ddos-that-knocked-spamhaus-offline-and-ho">Permalink</a> 

	| <a href="http://blog.cloudflare.com/the-ddos-that-knocked-spamhaus-offline-and-ho#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="185" width="581" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-20/lwhhFndFymGjBljDbgFgkFHdHxtDanAdlqGlyqeHqAAmFJgkqIbEHsoIfsmc/spamhaus_ddos_attack.png">
        <media:thumbnail height="159" width="500" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-20/lwhhFndFymGjBljDbgFgkFHdHxtDanAdlqGlyqeHqAAmFJgkqIbEHsoIfsmc/spamhaus_ddos_attack.png.scaled500.png"/>
      </media:content>
      <media:content type="image/jpeg" height="159" width="305" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-03-20/oqehwirAwbrxpqdJBJenzxdmAefduBHaGCsdwBGjbEEEAlzqablaIuECumpB/spamhaus_logo.jpg">
        <media:thumbnail height="159" width="305" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-20/oqehwirAwbrxpqdJBJenzxdmAefduBHaGCsdwBGjbEEEAlzqablaIuECumpB/spamhaus_logo.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/jpeg" height="1322" width="3100" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-03-20/rnFosdBxIJCjCexwHzDthcAawvbeqwtqFkyrImksJkqdgBrDwGAhGskhGadm/im_under_attack.jpg">
        <media:thumbnail height="213" width="500" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-20/rnFosdBxIJCjCexwHzDthcAawvbeqwtqFkyrImksJkqdgBrDwGAhGskhGadm/im_under_attack.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/jpeg" height="259" width="400" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-03-20/IerauHyfaagukmpyqyfBqClrfybyjBEHmrBknaJfpGlmItuGbfakiylsCfxy/burst_pipe.jpg">
        <media:thumbnail height="259" width="400" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-20/IerauHyfaagukmpyqyfBqClrfybyjBEHmrBknaJfpGlmItuGbfakiylsCfxy/burst_pipe.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/png" height="131" width="481" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-03-20/snzgfjctdrklJyqvsunggCHefAIHfqbIpCHeuIjCcvibdgujIHDsbIfAkFeu/history_repeats_itself.png">
        <media:thumbnail height="131" width="481" url="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-03-20/snzgfjctdrklJyqvsunggCHefAIHfqbIpCHeuIjCcvibdgujIHDsbIfAkFeu/history_repeats_itself.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Wed, 13 Mar 2013 09:52:00 -0700</pubDate>
      <title>CloudFlare Keeps TheBayLights.org Running Bright</title>
      <link>http://blog.cloudflare.com/cloudflare-keeps-thebaylightsorg-running-brig</link>
      <guid>http://blog.cloudflare.com/cloudflare-keeps-thebaylightsorg-running-brig</guid>
      <description>
        <![CDATA[<p>
	<p class="p1"><div class='p_embed p_image_embed'>
<a href="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-03-13/rxxonuzaxdwqAgjalhpzFJqxBAqaIrghiDzDunphzvkCFoujwBuFiAzHGBdu/bridge1.jpeg.scaled1000.jpg"><img alt="Bridge1" height="251" src="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-13/rxxonuzaxdwqAgjalhpzFJqxBAqaIrghiDzDunphzvkCFoujwBuFiAzHGBdu/bridge1.jpeg.scaled500.jpg" width="500" /></a>
</div>
</p>
<p class="p1"><span style="font-size: small;"><strong>The Art</strong></span></p>
<p class="p3"><span style="font-size: small;">When you think of San Francisco, undoubtedly one bridge in particular comes to mind - The Golden Gate Bridge. This year, however, the Bay Bridge is getting its moment in the spotlight thanks to </span><a href="http://wordspicturesideas.com/"><span class="s1">Words Pictures Ideas</span></a>, a CloudFlare customer.</p>
<p class="p4"><span class="s2">Words Pictures Ideas ser</span>vices brands and organizations in need of smarter communications. While thinking of ways to commemorate the 75th anniversary of the Bay Bridge, WPI founder Ben Davis came up with the <a href="http://wordspicturesideas.com/projects/illuminating-the-bay/"><span class="s1">idea to turn the West Span of the bridge into a canvas for light art</span></a>.</p>
<p class="p4"><span class="s2">Partnering with </span>internationally renowned light artist<span class="s2"> Leo Villareal, Ben and the WPI team began working on what would </span>become the world&rsquo;s largest LED light sculpture: <a href="http://thebaylights.org/"><span class="s3">The Bay Lights</span></a>.<span class="s2">&nbsp;</span></p>
<p class="p2"><strong style="font-size: small;"><br />The Plan</strong></p>
<p class="p3">The Bay Lights were officially unveiled on March 5, 2013. Brian VanderZanden, Lead Developer at WPI, knew there would be a surge in traffic to TheBayLights.org leading up to that day, and most likely a huge surge in traffic on the day of the unveiling. WPI has many sites on CloudFlare, including TheBayLights.org. He reached out to CloudFlare to make sure the site was ready to handle the increase in traffic.</p>
<p class="p3">CloudFlare suggested a few small optimizations (minification, an image that wasn't proxied because on a "grey cloud" DNS record), one useful reminder (<a href="https://support.cloudflare.com/entries/22055137-why-do-my-server-logs-show-cloudflare-s-ips-using-cloudflare"><span class="s1">whitelist</span></a> the <a href="http://www.cloudflare.com/ips"><span class="s1">CloudFlare IPs</span></a>), and a powerful recommendation: <a href="http://blog.cloudflare.com/introducing-pagerules-advanced-caching" target="_blank">Cache Everything</a>.</p>
<p class="p5">By default, CloudFlare will cache <a href="https://support.cloudflare.com/entries/22037282-what-file-extensions-does-cloudflare-cache-for-static-content"><span class="s1">obviously static assets</span></a>, but pass dynamic HTML through to the customer's webserver. For heavy load, on content that is not changing rapidly, full HTML pages -- or the entire site -- can be delivered from CloudFlare's global network, preserving the customer's webserver, database and other infrastructure. (Note: combined with <a href="http://blog.cloudflare.com/introducing-single-file-purge"><span class="s1">single file purge</span></a>, CloudFlare can serve as the global network for delivering a static site even with rapid changes, much as the <a href="http://kylerush.net/blog/meet-the-obama-campaigns-250-million-fundraising-platform/"><span class="s1">Obama 2012 web team did</span></a>.)</p>
<p class="p5"><br /><strong>The Results</strong></p>
<p class="p2">On the day of the unveiling, with Cache Everything turned on, TheBayLights.org saw traffic increase with a decrease on their system&rsquo;s resource utilization.</p>
<p class="p2"><div class='p_embed p_image_embed'>
<a href="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-03-13/thejitiFvAdcoaotmbvDFEJHHHicnlgqtzqAciEEEeJdeExzfqFgIcvugkqD/BayLights1.jpeg.scaled1000.jpg"><img alt="Baylights1" height="341" src="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-03-13/thejitiFvAdcoaotmbvDFEJHHHicnlgqtzqAciEEEeJdeExzfqFgIcvugkqD/BayLights1.jpeg.scaled500.jpg" width="500" /></a>
</div>
</p>
<p class="p6">By mid-day, a rush in traffic caused more load than the event's peak at 8:00 pm. The graph below shows an interesting resource demand for the site pre/post cache everything:</p>
<p class="p6"><div class='p_embed p_image_embed'>
<a href="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-03-13/gsJhHdbxiwevpxmlshoyxljzErvIfkAknzFcnijCwiFtsalcFrAwyhDEnqgl/cpu_vs_event.png.scaled1000.png"><img alt="Cpu_vs_event" height="177" src="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-03-13/gsJhHdbxiwevpxmlshoyxljzErvIfkAknzFcnijCwiFtsalcFrAwyhDEnqgl/cpu_vs_event.png.scaled500.png" width="500" /></a>
</div>
</p>
<p class="p6">The site saw the largest influx of traffic between 8:00-9:00 pm, but the average I/O during that hour was under 2Mb/s. By midnight traffic was back down to only 2X of baseline traffic levels.&nbsp;</p>
<p class="p2"><div class='p_embed p_image_embed'>
<a href="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-13/ouswBtHjCpwEeFJInvgatrFywoziHubvhzbmqoCmxyabJzaniBFiAxeDultI/GoD_analytics_visits_by_hour.png.scaled1000.png"><img alt="God_analytics_visits_by_hour" height="190" src="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-03-13/ouswBtHjCpwEeFJInvgatrFywoziHubvhzbmqoCmxyabJzaniBFiAxeDultI/GoD_analytics_visits_by_hour.png.scaled500.png" width="500" /></a>
</div>
</p>
<p class="p6">&ldquo;We began to celebrate at 9:15 pm as we were confident that the peak in site traffic had been reached and there were no issues,&rdquo; said Brian. &ldquo;We are thrilled with the guidance and help CloudFlare offered in keeping us online during our biggest moment of the year, as well as the day to day performance and security they provide for all of our sites.&rdquo;</p>
<p class="p2">The Bay Lights will continue to shine for the next two years, creating yet another tourist stop in San Francisco. At CloudFlare, we are excited to be a part of the experience and look forward to helping keep TheBayLights.org shining online.&nbsp;</p>
	
</p>

<p><a href="http://blog.cloudflare.com/cloudflare-keeps-thebaylightsorg-running-brig">Permalink</a> 

	| <a href="http://blog.cloudflare.com/cloudflare-keeps-thebaylightsorg-running-brig#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/1122929/CloudFlare4-16.jpeg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/he6rIlp3ceWH8</posterous:profileUrl>
        <posterous:firstName>Kristin</posterous:firstName>
        <posterous:lastName>Tarr</posterous:lastName>
        <posterous:nickName>ktarr</posterous:nickName>
        <posterous:displayName>Kristin Tarr</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="870" width="2355" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-13/mxskqsHzGcybtAzzcuAkqJpkbCwfqClGFlsczqfCyGvkgIcxiHjncyGkCwBF/GoD_5pm_satori.png">
        <media:thumbnail height="185" width="500" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-13/mxskqsHzGcybtAzzcuAkqJpkbCwfqClGFlsczqfCyGvkgIcxiHjncyGkCwBF/GoD_5pm_satori.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="240" width="677" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-03-13/gsJhHdbxiwevpxmlshoyxljzErvIfkAknzFcnijCwiFtsalcFrAwyhDEnqgl/cpu_vs_event.png">
        <media:thumbnail height="177" width="500" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-03-13/gsJhHdbxiwevpxmlshoyxljzErvIfkAknzFcnijCwiFtsalcFrAwyhDEnqgl/cpu_vs_event.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="394" width="1036" url="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-03-13/ouswBtHjCpwEeFJInvgatrFywoziHubvhzbmqoCmxyabJzaniBFiAxeDultI/GoD_analytics_visits_by_hour.png">
        <media:thumbnail height="190" width="500" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-03-13/ouswBtHjCpwEeFJInvgatrFywoziHubvhzbmqoCmxyabJzaniBFiAxeDultI/GoD_analytics_visits_by_hour.png.scaled500.png"/>
      </media:content>
      <media:content type="image/jpeg" height="321" width="640" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-03-13/rxxonuzaxdwqAgjalhpzFJqxBAqaIrghiDzDunphzvkCFoujwBuFiAzHGBdu/bridge1.jpeg">
        <media:thumbnail height="251" width="500" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-13/rxxonuzaxdwqAgjalhpzFJqxBAqaIrghiDzDunphzvkCFoujwBuFiAzHGBdu/bridge1.jpeg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/jpeg" height="865" width="1268" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-13/thejitiFvAdcoaotmbvDFEJHHHicnlgqtzqAciEEEeJdeExzfqFgIcvugkqD/BayLights1.jpeg">
        <media:thumbnail height="341" width="500" url="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-03-13/thejitiFvAdcoaotmbvDFEJHHHicnlgqtzqAciEEEeJdeExzfqFgIcvugkqD/BayLights1.jpeg.scaled500.jpg"/>
      </media:content>
    </item>
    <item>
      <pubDate>Thu, 07 Mar 2013 08:45:00 -0800</pubDate>
      <title>Go London User Group</title>
      <link>http://blog.cloudflare.com/go-london-user-group</link>
      <guid>http://blog.cloudflare.com/go-london-user-group</guid>
      <description>
        <![CDATA[<p>
	<p>We've mentioned before that we're using <a href="http://blog.cloudflare.com/go-at-cloudflare">Go</a> internally for projects such as Railgun (and a new DNS server and SSL infrastructure amongst other&nbsp;things). And we've mentioned that we are <a href="http://blog.cloudflare.com/cloudflare-london-were-hiring">opening an office in London</a>.</p>
<p>Now we're putting those two things together by sponsoring and helping to organize a new <a href="http://www.meetup.com/Go-London-User-Group/">Go London User Group</a>. The first meeting is on <a href="http://www.meetup.com/Go-London-User-Group/events/108066652/">March 27</a>&nbsp;at Makers Academy.&nbsp;CloudFlare will be providing food and drink. Speakers will be announced closer to the date. Be sure to sign up as we have limited space (and if it's full please put yourself on the waiting list so we can gauge how large the interest is).</p>
<p>Feel free to suggest speakers and talks of interest in the comments.</p>
<p>We're also actively <a href="http://blog.cloudflare.com/do-you-want-to-work-with-go">hiring Go</a> (and other) programmers in London and San Francisco.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/go-london-user-group">Permalink</a> 

	| <a href="http://blog.cloudflare.com/go-london-user-group#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/1943163/jgc.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/lAAjkS0LMSs0G</posterous:profileUrl>
        <posterous:firstName>John </posterous:firstName>
        <posterous:lastName>Graham-Cumming</posterous:lastName>
        <posterous:nickName>jgrahamc</posterous:nickName>
        <posterous:displayName>John  Graham-Cumming</posterous:displayName>
      </posterous:author>
    </item>
    <item>
      <pubDate>Wed, 06 Mar 2013 09:43:00 -0800</pubDate>
      <title>Load Balancing without Load Balancers</title>
      <link>http://blog.cloudflare.com/cloudflares-architecture-eliminating-single-p</link>
      <guid>http://blog.cloudflare.com/cloudflares-architecture-eliminating-single-p</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<a href="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-03-05/DmJzGGGyppFjjFGJoGEkffJEkppnuadqDiwfJslCwffdnpFEqasgqfAdsGDt/balance.jpg.scaled1000.jpg"><img alt="Balance" height="375" src="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-03-05/DmJzGGGyppFjjFGJoGEkffJEkppnuadqDiwfJslCwffdnpFEqasgqfAdsGDt/balance.jpg.scaled500.jpg" width="500" /></a>
</div>
</p>
<p>CloudFlare had an <a href="http://blog.cloudflare.com/todays-outage-post-mortem-82515" target="_blank">hour-long outage</a> this last weekend. Thankfully, outages like this have been a relatively rare occurance for our service. This is in spite of hundreds of thousands of customers, the enormous volume of legitimate traffic they generate, and the barrage of large denial of service attacks we are constantly <a href="http://blog.cloudflare.com/65gbps-ddos-no-problem" target="_blank">mitigating on their behalf</a>. While last weekend's outage exposed a flaw in our architecture that we're working to fully eliminate, largely our systems have been designed to be balanced and have no single points of failure. We haven't talked much about the architecture of CloudFlare's systems but thought the rest of the community might benefit from seeing some of the choices we've made, how we load balance our systems, and how this has allowed us to scale quickly and efficiently.</p>
<p><strong>Failure Isn't an Option, It's a Fact</strong></p>
<p>CloudFlare's architecture starts with an assumption: failure is going to happen. As a result, we have to plan for failure at every level and design a system that gracefully handles it when it occurs. To understand how we do this, you have to understand the components of CloudFlare's edge systems. Here are four critical components we deploy at the edge of our network:</p>
<ol>
<li><strong>Network: </strong>CloudFlare's <a href="http://www.cloudflare.com/network-map" target="_blank">23 data centers</a> (internally we refer to them as PoPs) are connected to the rest of the world via multiple providers. These connections are both through transit (bandwidth) providers as well as other networks we directly peer with.</li>
<li><strong>Router: </strong>at the edge of each of our PoPs is a router. This router announces the paths packets take to CloudFlare's network from the rest of the Internet.</li>
<li><strong>Switch:</strong> within each PoP there will be one or more switches that aggregate traffic within the PoP's local area network (LAN).</li>
<li><strong>Server:</strong> behind each switch there are a collection of servers. These servers perform some of the key tasks to power CloudFlare's service including DNS resolution, proxying, caching, and logging.</li>
</ol>
<p>Those are the four components you'll find in the racks that we run in locations around the world. You'll notice some things from a typical hardware stack seem to be missing. For example, there's no hardware load balancer. The problem with hardware load balancers (and hardware firewalls, for that matter) is that they often become the bottleneck and create a single point of failure themselves. Instead of relying on a piece of hardware to load balance across our network, we use routing protocols to spread traffic and handle failure.</p>
<p><strong>Anycast Is Your Friend</strong></p>
<p>For most of the Internet, IP addressess correspond to a single device connected to the public Internet. In your home or office, you may have multiple devices sitting behind a gateway using network address translation (NAT), but there is only one public IP address and all the devices that sit behind the network use a unique private IP address (e.g., in the space 192.168.X.X or 10.X.X.X). The general rule on the Internet is one unique IP per devices. This is a routing scheme known as Unicast. However, it's not the only way.</p>
<p>There are four major routing schemes: Unicast, Multicast, Broadcast, and Anycast. Multicast and Broadcast are so-called one-to-many routing schemes. With Broadcast, one node sends packets that hit all recipient nodes. Broadcast is not widely used any longer and was actually not implemented in IPv6 (its largest contemporary use has likely been <a href="http://blog.cloudflare.com/deep-inside-a-dns-amplification-ddos-attack" target="_blank">launching SMURF DDoS attacks</a>). With Multicast, one node sends packets that hit multiple (but not all) recipient nodes that have opted into a group (e.g., how a cable company may deliver a television broadcast over an IP network).</p>
<p><div class='p_embed p_image_embed'>
<img alt="Anycast" height="167" src="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-05/bEGigbozFGHHCwHulJwjssBJFfuhtgkHnJaIBbxfBqdtxoDxmHxsyCAliacr/anycast.png.scaled500.png" width="250" />
</div>
</p>
<p>Unicast and Anycast are one-to-one routing schemes. In both, there is one sender and one recipient of the packet. The difference between the two is that while there is only one possible destination on the entire network for a packet sent over Unicast, with Anycast there are multiple possible destinations and the network itself picks the route that is most preferential. On the wide area network (WAN) -- aka. the Internet -- this preference is for the shortest path from the sender to the recipient. On the LAN, the preferences can be set with weights that are honored by the router.</p>
<p><strong>Anycast at the WAN</strong></p>
<p>At CloudFlare, we use Anycast at two levels: the WAN and the LAN. At the WAN level, every router in all of CloudFlare's 23 data centers announces all of our external-facing IP addresses. For example, one of the IPs that CloudFlare announces for DNS services is&nbsp;173.245.58.205. A route to that IP address is announced from all 23 CloudFlare data centers. When you send a packet to that IP address, it passes through a series of routers. Those routers look at the available paths to CloudFlare's end points and send the packet down the one with the fewest stops along the way (i.e., "hops"). You can run a traceroute to see each of these steps.</p>
<p>If I run a traceroute from CloudFlare's office in San Francisco, the path my packets take is:</p>
<div class="CodeRay">
  <div class="code"><pre>$ traceroute 173.245.58.205
traceroute to 173.245.58.205 (173.245.58.205), 64 hops max, 52 byte packets
1  192.168.2.1 (192.168.2.1)  3.473 ms  1.399 ms  1.247 ms
2  10.10.11.1 (10.10.11.1)  3.136 ms  2.857 ms  3.206 ms
3  ge-0-2-5.cr1.sfo1.us.nlayer.net (69.22.X.X)  2.936 ms  3.405 ms  3.193 ms
4  ae3-70g.cr1.pao1.us.nlayer.net (69.22.143.170)  3.638 ms  4.076 ms  3.911 ms
5  ae1-70g.cr1.sjc1.us.nlayer.net (69.22.143.165)  4.833 ms  4.874 ms  4.973 ms
6  ae1-40g.ar2.sjc1.us.nlayer.net (69.22.143.118)  8.926 ms  8.529 ms  6.742 ms
7  as13335.xe-8-0-5.ar2.sjc1.us.nlayer.net (69.22.130.146)  5.048 ms
8  173.245.58.205 (173.245.58.205)  4.601 ms  4.338 ms  4.611 ms</pre></div>
</div>

<p>If you run the same traceroute from a Linode server in London, the path my packets take is:</p>
<div class="CodeRay">
  <div class="code"><pre>$ traceroute 173.245.58.205
traceroute to 173.245.58.205 (173.245.58.205), 30 hops max, 60 byte packets
1  212.111.X.X (212.111.X.X)  6.574 ms  6.514 ms  6.522 ms
2  212.111.33.X (212.111.33.X)  0.934 ms  0.935 ms  0.969 ms
3  85.90.238.69 (85.90.238.69)  1.396 ms  1.381 ms  1.405 ms
4  ldn-b3-link.telia.net (80.239.167.93)  0.700 ms  0.696 ms  0.670 ms
5  ldn-bb1-link.telia.net (80.91.247.24)  2.349 ms  0.700 ms  0.671 ms
6  ldn-b5-link.telia.net (80.91.246.147)  0.759 ms  0.771 ms  0.774 ms
7  cloudflare-ic-154357-ldn-b5.c.telia.net (80.239.161.246)  0.917 ms  0.853 ms  0.833 ms
8  173.245.58.205 (173.245.58.205)  0.972 ms  1.292 ms  0.916 ms</pre></div>
</div>

<p>In both cases, the 8th and final hop is the same. You can tell, however, that they are hitting different CloudFlare data centers from hints in the 7th hop (highlighted in red below): as13335.xe-8-0-5.ar2.<strong><span style="color: #ff0000;">sjc</span></strong>1.us.nlayer.net suggesting it is hitting San Jose and cloudflare-ic-154357-<strong><span style="color: #ff0000;">ldn</span></strong>-b5.c.telia.net suggesting it is hitting London.</p>
<p>Since packets will follow the shortest path, if a particular path is withdrawn then packets will find their way to the next shortest available route. For simple protocols like UDP that don't maintain state, Anycast is ideal and it has been used widely to load balance DNS for some time. At CloudFlare, we've done a significant amount of engineering to allow TCP to run across Anycast without flapping. This involves carefully adjusting routes in order to get optimal routing and also adjusting the way we handle protocol negotiation itself. While more complex to maintain than a Unicast network, the benefit is we can lose an entire data center and packets flow to the next closest facility without anyone noticing and hiccup.</p>
<p><strong>Anycast in the LAN</strong></p>
<p>Once a packet arrives as a particular CloudFlare data center we want to ensure it gets to a server that can correctly handle the request. There are four key tasks that CloudFlare's servers perform: DNS, proxy, cache, and logging. We tend to follow the Google-like approach and deploy generic, white-box servers that can perform a number of different functions. (Incidentally, if anyone is interested, we're thinking of doing a blog post to "tour" a typical CloudFlare server and discuss the choices we made in working with manufacturers to design them.) Since servers can fail or be overloaded, we need to be able to route traffic intelligently around problems. For that, we return to our old friend Anycast.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Bird" height="183" src="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-03-05/bADucsFgFzesjopHJztDmBaiIfwlleuoddxICCAIvlHfzArdjFdBwqcvhhkd/bird.jpeg.scaled500.jpg" width="275" />
</div>
</p>
<p>Using Anycast, each server within each of CloudFlare's data centers is setup to receive traffic from any of our public IP addresses. The routes to these servers are announced via the border gateway protocol (BGP) from the servers themselves. To do this we use a piece of software called <a href="http://bird.network.cz/" target="_blank">Bird</a>. (You can tell it's an awesomely intense piece of networking software just by looking at <a href="http://artax.karlin.mff.cuni.cz/~zajio1am/" target="_blank">one of its developers</a>.) While all servers announce a route across the LAN for all the IPs, each server assigns its own weight to each IPs route. The router is then configured such that the route with the lowest weight is preferred.</p>
<p>If a server crashes, Bird stops announcing the BGP route to the router. The router then begins sending traffic to the server with the next-lowest weighted route. We also monitor critical processes on each server. If any of these critical processes fails then it can signal Bird to withdraw a route. This is not all or nothing. The monitor is aware of the server's own load as well as the load on the other servers in the data center. If a particular server starts to become overloaded, and it appears there is sufficient capacity elsewhere, then just some of the BGP routes can be withdrawn to take some traffic away from the overloaded server.</p>
<p>Beyond failover, we are beginning to experiment with BGP to do true load balancing. In this case, the weights to multiple servers are the same and the router hashes the source IP, destination IP, and port in order to consistently route traffic to the same server. The hash mapping table can be adjusted to increase or decrease load to any machine in the cluster. This is relatively easy with simple protocols like UDP, so we're playing with it for DNS. It's trickier with protocols that need to maintain some session state, like TCP, and gets trickier still when you throw in SSL, but we have some cool things in our lab that will help us better spread load across all the available resources.</p>
<p><strong>Failure Scenarios</strong></p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-03-05/wtFEJHGzGrblpAFgkmobrabgkEyHmukyskJqzjhfqclgkJeqbsqiaxjcpwfB/global_thermonuclear_war.jpg.scaled1000.jpg"><img alt="Global_thermonuclear_war" height="274" src="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-03-05/wtFEJHGzGrblpAFgkmobrabgkEyHmukyskJqzjhfqclgkJeqbsqiaxjcpwfB/global_thermonuclear_war.jpg.scaled500.jpg" width="500" /></a>
</div>
</p>
<p>To understand this architecture, it's useful to think through some common failure scenarios.</p>
<ol>
<li><strong>Process Crash:</strong> if a core process (DNS, proxy, cache, or logging) crashes then the monitor daemon running on the server detects the failure. The monitor signals Bird to withdraw the BGP routes that are routed to that process (e.g., if just DNS crashes then the IPs that are used for CloudFlare name servers will be withdrawn, but the server will still respond to proxy traffic). With the routes withdrawn, the router in the data center sends traffic to the route with the next-lowest weight. The monitor daemon restarts the DNS server and, after verifying it has come up cleanly, signals Bird to start announcing routes again.</li>
<li><strong>Server Crash:</strong> if a whole server crashes, Bird crashes along with it. All BGP routes to the server are withdrawn and the router sends traffic to the servers with the next lowest route weights. A monitor process on a control server within the data center attempts to reboot the box using the IPMI management interface and, if that fails, a power cycle from the fancy power strip (PDU). After the monitor process has verified the box has come back up cleanly, Bird is restarted and routes to the server are reinitiated.</li>
<li><strong>Switch Crash:</strong> if a switch fails, all BGP routes to the servers behind the switch are automatically withdrawn. The routers are configured if they lose sufficient routes to the machines to drop the IPs that correspond to those routes out of the WAN Anycast pool. Traffic fails over for those IPs to the next closest data center. Monitors both inside and outside the affected data center alert our networking team who monitor the network 24/7 that there has been a switch failure so they can investigate and attempt a reboot.</li>
<li><strong>Router Crash:</strong> if a router fails, all BGP routes across the WAN are withdrawn for the data center for which the router is responsible. Traffic to the data center automatically fails over to the next closest data center. Monitors both inside and outside the affected data center alert our networking team who monitor the network 24/7 that there has been a router failure so they can investigate and attempt a reboot.</li>
<li><strong>Global Thermonuclear War:</strong> would be bad, but CloudFlare may continue to be able to route traffic to whatever portion of the Internet is left. As facilities were vaporized (starting with Las Vegas) their routers would stop announcing routes. As long as some facilities remained connected to whatever remained of the nextwork (maybe Sydney, Australia?) they would provide a path for traffic destined for our customers. We've designed the network such that more than half of it can completely fail and we'll still be able to keep up with the traffic.</li>
</ol>
<p>It's a rare company our size that gets to play with systems to globally load balance Internet-scale traffic. While we've done a number of smart things to build a very fault tolerant network, last weekend's events prove there is more work to be done. If these are the sort of problems that excite you and you're interested in helping build a network that can survive almost anything, <a href="https://www.cloudflare.com/join-our-team" target="_blank">we're hiring</a>.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/cloudflares-architecture-eliminating-single-p">Permalink</a> 

	| <a href="http://blog.cloudflare.com/cloudflares-architecture-eliminating-single-p#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/jpeg" height="768" width="1024" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-05/DmJzGGGyppFjjFGJoGEkffJEkppnuadqDiwfJslCwffdnpFEqasgqfAdsGDt/balance.jpg">
        <media:thumbnail height="375" width="500" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-03-05/DmJzGGGyppFjjFGJoGEkffJEkppnuadqDiwfJslCwffdnpFEqasgqfAdsGDt/balance.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/png" height="167" width="250" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-03-05/bEGigbozFGHHCwHulJwjssBJFfuhtgkHnJaIBbxfBqdtxoDxmHxsyCAliacr/anycast.png">
        <media:thumbnail height="167" width="250" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-03-05/bEGigbozFGHHCwHulJwjssBJFfuhtgkHnJaIBbxfBqdtxoDxmHxsyCAliacr/anycast.png.scaled500.png"/>
      </media:content>
      <media:content type="image/jpeg" height="183" width="275" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-03-05/bADucsFgFzesjopHJztDmBaiIfwlleuoddxICCAIvlHfzArdjFdBwqcvhhkd/bird.jpeg">
        <media:thumbnail height="183" width="275" url="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-03-05/bADucsFgFzesjopHJztDmBaiIfwlleuoddxICCAIvlHfzArdjFdBwqcvhhkd/bird.jpeg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/jpeg" height="466" width="850" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-03-05/wtFEJHGzGrblpAFgkmobrabgkEyHmukyskJqzjhfqclgkJeqbsqiaxjcpwfB/global_thermonuclear_war.jpg">
        <media:thumbnail height="274" width="500" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-03-05/wtFEJHGzGrblpAFgkmobrabgkEyHmukyskJqzjhfqclgkJeqbsqiaxjcpwfB/global_thermonuclear_war.jpg.scaled500.jpg"/>
      </media:content>
    </item>
    <item>
      <pubDate>Sun, 03 Mar 2013 05:47:00 -0800</pubDate>
      <title>Today's Outage Post Mortem</title>
      <link>http://blog.cloudflare.com/todays-outage-post-mortem-82515</link>
      <guid>http://blog.cloudflare.com/todays-outage-post-mortem-82515</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<a href="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-03-03/BpjbDjzGnokgkhdDeDoJqpbJAnusgrtvlwqbHvEErwsGAfmjobrkrEIrtsam/cloudflare_outage.png.scaled1000.png"><img alt="Cloudflare_outage" height="151" src="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-03-03/BpjbDjzGnokgkhdDeDoJqpbJAnusgrtvlwqbHvEErwsGAfmjobrkrEIrtsam/cloudflare_outage.png.scaled500.png" width="500" /></a>
</div>
</p>
<p>This morning at&nbsp;<span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">09:47 UTC CloudFlare effectively dropped off the Internet. The outage affected all of CloudFlare's services including DNS and any services that rely on our web proxy. During the outage, anyone accessing CloudFlare.com or any site on CloudFlare's network would have received a DNS error. Pings and Traceroutes to CloudFlare's network resulted in a "No Route to Host" error.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">The cause of the outage was a system-wide failure of our edge routers. CloudFlare currently runs 23 data centers worldwide. These data centers are connected to the rest of the Internet using routers. These routers announce the path that, from any point on the Internet,&nbsp;</span><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">packets</span><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">&nbsp;should use to reach our network. When a router goes down, the routes to the network that sits behind the router are withdrawn from the rest of the Internet.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">We regularly will shut down one or a small handful of routers when we are upgrading a facility. Because we use Anycast, traffic naturally fails to the next closest data center. However, this morning we encountered a bug that caused all of our routers to fail network wide.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;"><strong>Flowspec</strong></span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">We are largely a Juniper shop at CloudFlare and all the edge routers that were affected were from Juniper. One of the reasons we like Juniper is their support of a&nbsp;<a href="http://www.slideshare.net/sfouant/an-introduction-to-bgp-flow-spec" target="_blank">protocol called Flowspec</a>. Flowspec allows you to propagate router rules to a large number of routers efficiently.&nbsp;</span><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">At CloudFlare, we constantly make updates to the rules on our routers. We do this to fight attacks as well as to shift traffic so it can be served as fast as possible.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">This morning, we saw a DDoS attack being launched against one of our customers. The attack specifically targeted the customer's DNS servers. We have an internal tool that profiles attacks and outputs signatures that our automated systems as well as our ops team can use to stop attacks. Often, we use these signatures in order to create router rules to either rate limit or drop known-bad requests.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">In this case, our attack profiler output the fact that the attack packets were between&nbsp;</span><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">99,971 and 99,985 bytes long. That's odd to begin with because the largest packets sent across the Internet are typically in the 1,500-byte range and average around 500 &ndash; 600 bytes. We have the maximum packet size set to 4,470 on our network, which is on the large size, but well under what the attack profiler was telling us was the size of these attack packets.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;"><strong>Bad Rule</strong></span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">Someone from our operations team is monitoring our network 24/7. As is normal practice for us, one of our ops team members took the output from the profiler and added a rule based on its output to drop packets that were between 99,971 and 99,985 bytes long. Here's what the rule (somewhat simplified and with the IPs obscured) looked like in Junos, the Juniper operating system:</span></p>
<div class="CodeRay">
  <div class="code"><pre>+    route 173.X.X.X/32-DNS-DROP {
+        match {
+            destination 173.X.X.X/32;
+            port 53;
+            packet-length [ 99971 99985 ];
+        }
+        then discard;
+    }</pre></div>
</div>

<p>Flowspec accepted the rule and relayed it to our edge network. What should have happened is that no packet should have matched that rule because no packet was actually that large. What happened instead is that the routers encountered the rule and then proceeded to consume all their RAM until they crashed.</p>
<p>In all cases, we run a monitoring process that reboots the routers automatically when they crash. That worked in a few cases. Unfortunately, many of the routers crashed in such a way that they did not reboot automatically and we were not able to access the routers' management ports. Even though some data centers came back online initially, they fell back over again because all the traffic across our entire network hit them and overloaded their resources.</p>
<p><a href="https://twitter.com/sambowne" target="_blank">Sam Bowne</a>, a computer science professor at City College of San Francisco, used BGPlay to capture the following video of BGP sessions being withdrawn as our routers crashed:</p>
<p><iframe src="http://www.youtube.com/embed/wMRaKtydILI" frameborder="0" height="315" width="420"></iframe></p>
<p><strong><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">Incident Response</span></strong></p>
<p>CloudFlare's ops and network teams were aware of the incident immediately because of both internal and external monitors we run on our network. While it wasn't initially clear the reason the routers had crashed, it was clear that it was an issue caused by an inability for packets to find a route to our network. We were able to access some routers and see that they were crashing when they encountered this bad rule. We removed the rule and then called the network operations teams in the data centers where our routers were unresponsive to ask them to physically access the routers and perform a hard reboot.</p>
<p>CloudFlare's 23 data centers span 14 countries so the response took some time but within about 30 minutes we began to restore CloudFlare's network and services. By&nbsp;<span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">10:49 UTC, all of CloudFlare's services were restored. We continue to investigate some edge cases where people are seeing outages. In nearly all of these cases, the problem is that a bad DNS response has been cached. Typically clearing the DNS cache will resolve the issue.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">We have already reached out to Juniper to see if this is a known bug or something unique to our setup and the kind of traffic we were seeing at the time. We will be doing more extensive testing of Flowspec provisioned filters and evaluating whether there are ways we can isolate the application of the rules to only those data centers that need to be updated, rather than applying the rules network wide. Finally, we plan to proactively issue service credits to accounts covered by SLAs. Any amount of downtime is completely unacceptable to us and the whole CloudFlare team is sorry we let our customers down this morning.</span></p>
<p><strong><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">Parallels to Syria</span></strong></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">In writing this up, I was reminded of the parallels to the&nbsp;<a href="http://blog.cloudflare.com/how-syria-turned-off-the-internet" target="_blank">Syrian Internet outage</a>&nbsp;we reported on earlier this year. In that case, we were able to detect as the Syrian government shut down their board routers and effectively cut the country off from the rest of the Internet. In CloudFlare's case the cause was not intentional or malicious, but the net effect was the same: a router change caused a network to go offline.</span></p>
<p><span style="font-family: Helvetica Neue, Helvetica, Arial, Geneva, sans-serif;">At CloudFlare, we spend a significant amount of our time immersed in the dark arts of Internet routing. This incident, like the incident in Syria, illustrates the power and importance of the these network protocols. We let our customer down this morning, but we will learn from the incident and put more controls in place to eliminate problems like this in the futre.</span></p>
<p>&nbsp;</p>
	
</p>

<p><a href="http://blog.cloudflare.com/todays-outage-post-mortem-82515">Permalink</a> 

	| <a href="http://blog.cloudflare.com/todays-outage-post-mortem-82515#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="159" width="526" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-03-03/BpjbDjzGnokgkhdDeDoJqpbJAnusgrtvlwqbHvEErwsGAfmjobrkrEIrtsam/cloudflare_outage.png">
        <media:thumbnail height="151" width="500" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-03-03/BpjbDjzGnokgkhdDeDoJqpbJAnusgrtvlwqbHvEErwsGAfmjobrkrEIrtsam/cloudflare_outage.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Thu, 28 Feb 2013 09:47:00 -0800</pubDate>
      <title>CloudFlare London: We're hiring</title>
      <link>http://blog.cloudflare.com/cloudflare-london-were-hiring</link>
      <guid>http://blog.cloudflare.com/cloudflare-london-were-hiring</guid>
      <description>
        <![CDATA[<p>
	<p>When we talk about international expansion we're usually talking about adding data centers around the world. The last one we added was in <a href="http://blog.cloudflare.com/seoul-korea-cloudflares-23rd-data-center">Seoul, South Korea</a>. And we've had a data center in London for a very long time. But now we're adding something different: people.</p>
<p>As CloudFlare's customer base and network have grown our need for 24 hour operations and technical support has grown. At the moment keeping things running means keeping people awake in California. With data centers in <a href="http://www.cloudflare.com/network-map">23 locations</a> around the world and customers in every country CloudFlare staff have to keep things humming day and night.</p>
<p>And so CloudFlare will expand in the next couple of months with an office in London.</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-02-28/vsFkfddlvECwIaHckvmrBzxIupCxmedheCxBnwyJFgCeesJsomnnyuGAIAgc/_65027163_65027162.jpeg.scaled1000.jpg"><img alt="_65027163_65027162" height="281" src="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-02-28/vsFkfddlvECwIaHckvmrBzxIupCxmedheCxBnwyJFgCeesJsomnnyuGAIAgc/_65027163_65027162.jpeg.scaled500.jpg" width="500" /></a>
</div>
We believe that London will make an ideal base for operations and technical support to complement our San Francisco office, and we can dip into the rich London talent pool to find people.</p>
<p>We keep our&nbsp;<a href="http://www.cloudflare.com/join-our-team">Join Our Team</a> page updated with positions on in London. These include <a href="http://www.jobscore.com/jobs/cloudflare/technical-customer-support/cNW2NomN0r4QnGeJe4efaV?ref=rss&amp;sid=68">Technical Customer Support</a>&nbsp;and <a href="http://www.jobscore.com/jobs/cloudflare/technical-operations-engineer/bcCGgQkAmr4lnoeJe4bk1X?ref=rss&amp;sid=68">Technical Operations Engineer</a>.</p>
<p>Keep an eye out for new openings as we expand into London.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/cloudflare-london-were-hiring">Permalink</a> 

	| <a href="http://blog.cloudflare.com/cloudflare-london-were-hiring#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/1943163/jgc.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/lAAjkS0LMSs0G</posterous:profileUrl>
        <posterous:firstName>John </posterous:firstName>
        <posterous:lastName>Graham-Cumming</posterous:lastName>
        <posterous:nickName>jgrahamc</posterous:nickName>
        <posterous:displayName>John  Graham-Cumming</posterous:displayName>
      </posterous:author>
      <media:content type="image/jpeg" height="549" width="976" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-28/vsFkfddlvECwIaHckvmrBzxIupCxmedheCxBnwyJFgCeesJsomnnyuGAIAgc/_65027163_65027162.jpeg">
        <media:thumbnail height="281" width="500" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-02-28/vsFkfddlvECwIaHckvmrBzxIupCxmedheCxBnwyJFgCeesJsomnnyuGAIAgc/_65027163_65027162.jpeg.scaled500.jpg"/>
      </media:content>
    </item>
    <item>
      <pubDate>Wed, 27 Feb 2013 15:07:00 -0800</pubDate>
      <title>How to Tell How Well Railgun Is Working for Your Site</title>
      <link>http://blog.cloudflare.com/how-to-tell-how-well-railgun-is-working-for-y</link>
      <guid>http://blog.cloudflare.com/how-to-tell-how-well-railgun-is-working-for-y</guid>
      <description>
        <![CDATA[<p>
	<p>Yesterday, we announced that 30 of the world's largest hosting providers are now supporting CloudFlare's Railgun WAN optimization technology. Railgun uses delta compression to only transmit the parts of a dynamic page that have changed from one request to another. The net effect is that, on average, we can achieve a 99.6% compression ratio. In other words, what uncompressed would have taken 200 packets with Railgun we can transmit in a single packet.</p>
<p>This blog post is about using headers we include in responses delivered from Railgun in order to see how it is working. We've been running Railgun on CloudFlare.com for the last few months so I'll use it as an example.</p>
<p><strong>Exposing the Headers</strong></p>
<p>When a request is handled by Railgun, CloudFlare inserts a header with diagnostic information to track how the protocol is doing. If you want to see these headers, you'll need to use a browser that supports examining header information. Google's Chrome Browser or Apple's Safari Browser allow you to access Developer Tools (in Google, select the View &gt; Developer &gt; Developer Tools menu; in Safari, select the Develop &gt; Show Web Inspector menu). In Firefox, you can install<a href="http://getfirebug.com/" target="_blank">Firebug</a>&nbsp;to see response headers. Microsoft's Internet Explorer makes it a bit trickier to see the response headers, but you can use a tool like&nbsp;<a href="http://www.fiddler2.com/fiddler2/" target="_blank">Fiddler</a>&nbsp;in order to expose them.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Claire_screenshot" height="523" src="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-27/urmGxDihAsiwntpxntCseeHpiBGyoAoHAzqGmbnIJfFHJyzHdCgAGakbcetd/claire_screenshot.png.scaled500.png" width="264" />
</div>
</p>
<p>At CloudFlare, we've also made a Chrome extension for our own debugging purposes that we call Claire. When installed, it adds a small "cloud" icon to the right corner of the URL bar. If you're visiting a site that uses CloudFlare, lights up orange. Small icons under the cloud indicate whether you're using SPDY, Railgun, or IPv6 for your connection. Clicking on the icon exposes more data including information about the Railgun connection.</p>
<p>While Claire makes seeing the Railgun information easy, I'm going to walk through the rest of this post assuming you don't have it installed. Instead, I'll use Chrome's Developer Tools for the examples.</p>
<p><strong>Story in the Headers</strong></p>
<p>If you open the Developer Tools panel and click on the Network tab you'll see an interface like the one in the picture below:</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fihmvFBcupbBoDHnJsoschyDdsuAabJggrmHEpHnEDHqceglCxcviozBimfx/chrome_developer_tools_cloudflare.png.scaled1000.png"><img alt="Chrome_developer_tools_cloudflare" height="310" src="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fihmvFBcupbBoDHnJsoschyDdsuAabJggrmHEpHnEDHqceglCxcviozBimfx/chrome_developer_tools_cloudflare.png.scaled500.png" width="500" /></a>
</div>
</p>
<p>Clicking on the first item in the list, which represents the dynamic HTML content that makes up the page, and then clicking on the Headers tab will show you the headers your browser sent to CloudFlare's servers as well as, if you scroll down, the response headers that your browser received back. Below is a sample of the response headers when accessing <a href="http://www.cloudflare.com">www.cloudflare.com</a>:</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fIgArqCCpvyEbsjzhjtxcEiyxJysceHjkcdeqGkGyxAedHnrdvqdcFIoDyDz/cloudflare_response_headers.png.scaled1000.png"><img alt="Cloudflare_response_headers" height="167" src="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fIgArqCCpvyEbsjzhjtxcEiyxJysceHjkcdeqGkGyxAedHnrdvqdcFIoDyDz/cloudflare_response_headers.png.scaled500.png" width="500" /></a>
</div>
</p>
<p>There are two headers that CloudFlare is inserting in the response:</p>
<p style="margin: 0 0 0 0; padding: 0 0 0 30px;">cf-railgun: &nbsp; e95b1c46e0 0.02 0.037872 0030 9878</p>
<p style="margin: 0 0 0 0; padding: 0 0 0 30px;">cf-ray: &nbsp; 478149ad1570291</p>
<p>The second of these headers is what we call a RayID. This is a unique serial number attached to every request through the CloudFlare network, start to finish, which helps us diagnose if there's a problem at some step in our chain. If you ever have an error on your site when accessing CloudFlare, providing the RayID to our support team can help us track down the cause very quickly. The header I'm going to focus on for this post is the cf-railgun header, which I'll break down below.</p>
<p><strong>The CF-Railgun Header</strong></p>
<p>The CF-Railgun header has up to five codes separated by a space. In order, these codes and their corresponding values from the example above are:</p>
<ul>
<li>Railgun Request ID:&nbsp;e95b1c46e0</li>
<li>Compression Ratio: 0.02</li>
<li>Origin Processing Time: 0.037872</li>
<li>Railgun Flags: 0030</li>
<li>Version Number: 9878</li>
</ul>
<p>Breaking these down, the Railgun Request ID corresponds to an internal process number that allows us to track what connection handled a request in order to diagnose potential problems. Generally, you shouldn't need this value unless you're reporting a problem with your Railgun installation.</p>
<p>The Compression Ratio is more interesting in gauging how Railgun is down. It represents the size of the response after Railgun's delta compression expressed as a percentage. In the example above, the HTML returned for <a href="http://www.cloudflare.com">www.cloudflare.com</a> is 0.02% of the size of the original that would be returned assuming no origin compression. Another way of thinking about this is the amount of data saved, which can be calculated by subtracting the Compression Ratio value from 100. In this case, 99.98% of the data that would have been required to generate <a href="http://www.cloudflare.com">www.cloudflare.com</a> doesn't need to be transmitted because of the Railgun compression.</p>
<p>The Origin Processing Time represents the time, in seconds, that Railgun waits for the origin web server to generate the page. In this case, the origin server takes 0.03782 seconds from when the Railgun listener sends the request to the origin to when it responds. If this number is large, it means your web server or database may be hitting a bottleneck that is slowing down its time to render the full page.</p>
<p>The Railgun Flags represent how a request was processed. The simplified way of looking at the Railgun Flags is to see the 4-digit sequence as zzXz. Ignore the z's and focus on the number or letter in the X position. If it is 3,7, B or F then it means Railgun Compression is working correctly.</p>
<p>If there is an error of some sort, the Compression Ratio is likely to be listed as "normal" or "direct." This means that Railgun's compression was bypassed for one reason or another. The Railgun Flags help diagnose why. The Railgun Flags are a bitset and, in order to fully interpret them,, you need to use the rg-diag utility which is included with the&nbsp;<a href="https://www.cloudflare.com/resources-downloads" target="_blank">Railgun packages</a>. Run the utility from the command line with the -decode option. For example, to decode the Railgun Code 0038, for example, you'd run:</p>
<p style="padding-left: 30px;"><span style="color: #545454;"><span style=""><strong>rg-diag -decode=</strong></span></span><span style="color: #545454;"><span style=""><strong>0038</strong></span></span><span style="">&nbsp;</span></p>
<p>Which returns in:</p>
<p style="padding-left: 30px; margin: 0 0 0 0; padding-top: 0; padding-bottom: 0;"><span style="font-family: Menlo, monospace; font-size: xx-small;"><span style="">Railgun Flag Existing origin connection reused</span></span></p>
<p style="padding-left: 30px; margin: 0 0 0 0; padding-top: 0; padding-bottom: 0;"><span style="font-family: Menlo, monospace; font-size: xx-small;"><span style="">Railgun Flag rg-sender sent dictionary</span></span></p>
<p style="padding-left: 30px; margin: 0 0 0 0; padding-top: 0; padding-bottom: 0;"><span style="font-family: Menlo, monospace; font-size: xx-small;"><span style="">Railgun Flag rg-listener found dictionary</span></span></p>
<p>This information can be useful in diagnosing potential problems with Railgun. The good news is that the Railgun protocol is designed to be resilient. If a connection fails for some reason, in most cases it will immediately roll over to a normal HTTP or HTTPS connection without the visitor seeing an error.</p>
<p>Finally, returning to the cf-railgun header, the final variable is the Version Number which indicates the version of the Railgun Listener software that is running on the origin server's network. The numbers aren't necessarily sequential, so having a lower number than another Railgun Listener doesn't necessarily mean your Listener is out of date.</p>
<p><strong>Claire Makes It Easy</strong></p>
<p>The Claire Chrome Plugin simplifies the header, leaving out the Railgun Flags and Version Number. Instead, it returns the Railgun Request ID (useful to provide to our support team if there's an issue), the amount of data saved for the particular request (derived from 100 - the Compression Ratio), and the Origin Processing Time (in seconds). Generally, this is all you should need to see whether Railgun is functioning as intended on your site.</p>
<p>Stay tuned. We'll post more information on tips for getting the most out of Railgun, as well as some of the design and engineering considerations that went into designing the protocol, over the coming days.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/how-to-tell-how-well-railgun-is-working-for-y">Permalink</a> 

	| <a href="http://blog.cloudflare.com/how-to-tell-how-well-railgun-is-working-for-y#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="523" width="264" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-27/urmGxDihAsiwntpxntCseeHpiBGyoAoHAzqGmbnIJfFHJyzHdCgAGakbcetd/claire_screenshot.png">
        <media:thumbnail height="523" width="264" url="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-27/urmGxDihAsiwntpxntCseeHpiBGyoAoHAzqGmbnIJfFHJyzHdCgAGakbcetd/claire_screenshot.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="193" width="579" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fIgArqCCpvyEbsjzhjtxcEiyxJysceHjkcdeqGkGyxAedHnrdvqdcFIoDyDz/cloudflare_response_headers.png">
        <media:thumbnail height="167" width="500" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fIgArqCCpvyEbsjzhjtxcEiyxJysceHjkcdeqGkGyxAedHnrdvqdcFIoDyDz/cloudflare_response_headers.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="708" width="1143" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fihmvFBcupbBoDHnJsoschyDdsuAabJggrmHEpHnEDHqceglCxcviozBimfx/chrome_developer_tools_cloudflare.png">
        <media:thumbnail height="310" width="500" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-02-27/fihmvFBcupbBoDHnJsoschyDdsuAabJggrmHEpHnEDHqceglCxcviozBimfx/chrome_developer_tools_cloudflare.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Tue, 26 Feb 2013 00:51:00 -0800</pubDate>
      <title>CloudFlare's Railgun: Easier Than Ever</title>
      <link>http://blog.cloudflare.com/cloudflares-railgun-easier-than-ever</link>
      <guid>http://blog.cloudflare.com/cloudflares-railgun-easier-than-ever</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<a href="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-26/uHigIrcJDnIfmduglssaehtkJlurFqmzwxIenihcucHCEgwfgzeozEAJIffm/railgun.png.scaled1000.png"><img alt="Railgun" height="225" src="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-26/uHigIrcJDnIfmduglssaehtkJlurFqmzwxIenihcucHCEgwfgzeozEAJIffm/railgun.png.scaled500.png" width="500" /></a>
</div>
</p>
<p>Over the next few days we have a number of announcements regarding CloudFlare's Railgun technology. We wanted to begin, however, with what is in some ways the end: the ways in which you can take advantage of Railgun yourself. Today we're proud to announce two ways in which you can make your dynamic content faster than was ever possible before.</p>
<p><strong>Do It Yourself</strong></p>
<p>First, today CloudFlare is announcing version 3.3.3 of Railgun. This version has been battle tested on high-traffic sites including <a href="http://www.imgur.com/" target="_blank">Imgur</a> and <a href="http://www.4chan.org/" target="_blank">4chan</a>. It's run billions of requests in a number of different environments through the new protocol and we're ready to push it out to the world. To make the process of installing Railgun easy, we've released RPMs for most the popular Linux and BSD variants including:</p>
<ul>
<li>Ubuntu 12.10</li>
<li>Ubuntu 12.04</li>
<li>Ubuntu 11.10</li>
<li>Ubuntu 10.04 &nbsp;</li>
<li>FreeBSD 9</li>
<li>FreeBSD 8 &nbsp;</li>
<li>CentOS 6</li>
<li>CentOS 5 &nbsp;</li>
<li>Debian 6</li>
</ul>
<p>You can download any of these RPMs via the <a href="http://www.cloudflare.com/resources-downloads" target="_blank">CloudFlare Downloads page</a>. In addition, we've released an Amazon Machine Instance (AMI) for Amazon Web Services that you can install if you want to have you own Railgun listener. The AMI will be available soon via the AWS AMI manager.</p>
<p>The largest platform we're missing is Windows Server and we are working on updates to the Go runtime in order to allow us to compile for that platform and meet our quality standards. For the Windows Server users out there, stay tuned. We haven't forgotten about you.</p>
<p><strong>But Wait, There's More</strong></p>
<p>But that's not the really exciting part. We're extremely excited to announce that a majority of the world's leading hosting providers are now supporting CloudFlare's Railgun technology. These 30 hosting providers have already registered to be CloudFlare Optimized Hosts. That means you can enable Railgun, usually with a single click and without having to install any software or change any of your code. Within the next few days, all of the following hosts will be supporting CloudFlare's Railgun:</p>
<ul>
<li> <span style=""><a href="http://www.040hosting.eu/" target="_blank">040hosting</a></span></li>
<li><span style=""><a href="http://www.a2hosting.com/" target="_blank">A2 Hosting</a></span></li>
<li><span style=""><a href="http://www.arvixe.com/" target="_blank">Arvixe</a></span></li>
<li><span style=""><a href="http://www.bluehost.com/" target="_blank">Bluehost</a></span></li>
<li><span style=""><a href="http://byethost.com/" target="_blank">ByetHost</a></span></li>
<li><span style=""><a href="http://www.corecommerce.com/" target="_blank">CoreCommerce</a></span></li>
<li><span style=""><a href="http://dreamhost.com/" target="_blank">DreamHost</a></span></li>
<li><span style=""><a href="http://elserver.com/" target="_blank">ELServer</a></span></li>
<li><span style=""><a href="http://www.fastdomain.com/" target="_blank">FastDomain</a></span></li>
<li><span style=""><a href="http://www.greengeeks.com/" target="_blank">GreenGeeks</a></span></li>
<li><span style=""><a href="http://www.hostpapa.com/" target="_blank">HostPapa</a></span></li>
<li><span style=""><a href="http://www.hostmonster.com/" target="_blank">HostMonster</a></span></li>
<li><span style=""><a href="http://www.justhost.com/" target="_blank">Just Host</a></span></li>
<li><span style=""><a href="http://www.interserver.net/" target="_blank">InterServer</a></span></li>
<li><span style=""><a href="http://www.mapletime.com/" target="_blank">MapleTime</a></span></li>
<li><span style=""><a href="http://mediatemple.net/" target="_blank">(mt) Media Temple</a></span></li>
<li><span style=""><a href="http://www.mddhosting.com/" target="_blank">MDDHosting</a></span></li>
<li><span style=""><a href="http://www.namecheap.com/" target="_blank">NameCheap</a></span></li>
<li><span style=""><a href="http://www.pacifichost.com/" target="_blank">PacificHost</a></span></li>
<li><span style=""><a href="http://www.proisp.no/" target="_blank">PRO ISP</a></span></li>
<li><span style=""><a href="http://www.siteground.com/" target="_blank">SiteGround</a></span></li>
<li><span style=""><a href="http://www.sliqua.com/" target="_blank">Sliqua Enterprise Hosting</a></span></li>
<li><span style=""><a href="http://www.softcloud.co.uk/" target="_blank">Softcloud Hosting</a></span></li>
<li><span style=""><a href="https://www.sparkred.com/" target="_blank">SparkRed</a></span></li>
<li><span style=""><a href="https://ventraip.com.au/" target="_blank">VentraIP</a></span></li>
<li><span style=""><a href="http://vexxhost.com/" target="_blank">VEXXHOST</a></span></li>
<li><span style=""><a href="http://www.webhostingbuzz.com/" target="_blank">WebHostingBuzz</a></span></li>
<li><span style=""><a href="http://www.webhostingpad.com/" target="_blank">WebHostingPad</a></span></li>
<li><span style=""><a href="http://x10hosting.com/" target="_blank">x10Hosting</a></span></li>
<li><span style=""><a href="http://zuver.net.au/" target="_blank">Zuver</a></span></li>
</ul>
<p>In most cases, if you're already hosted with one of these hosting providers, getting the benefits of Railgun is free for you. If you're using one of these hosts, look for an option in your control panel to enable Railgun. If you're not already using one of these hosts, but you want to use Railgun, you should either contact your hosting provider to <a href="https://www.cloudflare.com/partner-programs" target="_blank">become a CloudFlare Optimized Partner</a>, or consider switching to one of the providers above.</p>
<p>Railgun is a revolutionary new protocol that makes dynamic web performance significantly faster and less bandwidth-intensive than was ever previously possible. Over the next few days, we'll be releasing more details about the protocol. In the meantime, we wanted to make sure you knew where you could go to get Railgun today if you're interested. Stay tuned for more.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/cloudflares-railgun-easier-than-ever">Permalink</a> 

	| <a href="http://blog.cloudflare.com/cloudflares-railgun-easier-than-ever#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="360" width="800" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-26/uHigIrcJDnIfmduglssaehtkJlurFqmzwxIenihcucHCEgwfgzeozEAJIffm/railgun.png">
        <media:thumbnail height="225" width="500" url="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-26/uHigIrcJDnIfmduglssaehtkJlurFqmzwxIenihcucHCEgwfgzeozEAJIffm/railgun.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Fri, 22 Feb 2013 13:12:00 -0800</pubDate>
      <title>Good Web Security News: Open DNS Resolvers Are Getting Closed</title>
      <link>http://blog.cloudflare.com/good-news-open-dns-resolvers-are-getting-clos</link>
      <guid>http://blog.cloudflare.com/good-news-open-dns-resolvers-are-getting-clos</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<img alt="Good_news" height="189" src="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-22/acDJcnqgEakaqABIJkcbuHuAAmxqerjoauDvluIjIGcEqunFgmaytJFCxikf/good_news.jpeg.scaled500.jpg" width="267" />
</div>
</p>
<p>This has been a rough week in the security industry with big attacks and compromises reported at companies from Facebook to Apple. We're therefore happy to end the week with some good news: the web's open resolvers, one of the sources of the biggest DDoS attacks, are getting closed.</p>
<p><strong>Sad State of Affairs</strong></p>
<p>Last October, we wrote a <a href="http://blog.cloudflare.com/deep-inside-a-dns-amplification-ddos-attack" target="_blank">blog post about DDoS amplification attacks</a>. This type of attack makes up some of the largest DDoSs CloudFlare sees, sometimes exceeding 100 gigabits per second (100Gbps). The attacks use DNS resolvers that haven't been properly secured in order to "amplify" the resources of the attacker. An attacker can achieve more than a 50x amplification, meaning that for every byte they are able to generate themselves they can pummel a victim with 50 bytes of garbage data.</p>
<p>The problem stems from misconfigured DNS resolver software (e.g., BIND) that is setup to respond to a query from any IP address. Since DNS requests typically are sent over UDP, which, unlike TCP, does not require a handshake, an attacker can spoof a victim's IP address as the source address in a packet and a misconfigured DNS resolver will happily bombard the victim with responses.</p>
<p><strong>Closing the Open Resolvers</strong></p>
<p>While CloudFlare's network is very good at absorbing even these large attacks, the long term solution for the web is for providers to clean up the open resolvers running on their networks. We wanted to help with that so we engaged in a bit of name-and-shame at the end of the last blog post, listing the networks with the largest number of open resolvers. The good news is it worked: almost four months later our tests show that the number of open resolvers across the Internet is down more than 30%. The chart below shows the progress individual networks have made in cleaning up the problem.</p>
<table border="0" style="border-collapse: collapse; font-size: 11px;" width="500">

<tr height="20">
<td class="xl30" height="16"><strong><span style="text-decoration: underline;">ASN</span></strong></td>
<td class="xl30"><span style="text-decoration: underline;"><strong>Network</strong></span></td>
<td class="xl30"><span style="text-decoration: underline;"><strong>10/30/12</strong></span></td>
<td class="xl30"><span style="text-decoration: underline;"><strong>2/22/13</strong></span></td>
<td class="xl30"><span style="text-decoration: underline;"><strong>% Change</strong></span></td>
</tr>
<tr height="20">
<td class="xl27" height="15">21844 &nbsp;</td>
<td class="xl28">THEPLANET-AS - ThePlanet.com Internet Services, In</td>
<td class="xl28" align="right">2925</td>
<td class="xl28" align="right">2216</td>
<td class="xl29" align="right">-24%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">3462 &nbsp; &nbsp;</td>
<td class="xl28">HINET Data Communication Business Group</td>
<td class="xl28" align="right">2739</td>
<td class="xl28" align="right">2213</td>
<td class="xl29" align="right">-19%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">36351 &nbsp;</td>
<td class="xl28">SOFTLAYER - SoftLayer Technologies Inc.</td>
<td class="xl28" align="right">1075</td>
<td class="xl28" align="right">781</td>
<td class="xl29" align="right">-27%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">9394 &nbsp; &nbsp;</td>
<td class="xl28">CRNET CHINA RAILWAY Internet(CRNET)</td>
<td class="xl28" align="right">1052</td>
<td class="xl28" align="right">774</td>
<td class="xl29" align="right">-26%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">4713 &nbsp; &nbsp;</td>
<td class="xl28">OCN NTT Communications Corporation</td>
<td class="xl28" align="right">1044</td>
<td class="xl28" align="right">722</td>
<td class="xl29" align="right">-31%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">45595 &nbsp;</td>
<td class="xl28">PKTELECOM-AS-PK Pakistan Telecom Company Limited</td>
<td class="xl28" align="right">1030</td>
<td class="xl28" align="right">716</td>
<td class="xl29" align="right">-30%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">4134 &nbsp; &nbsp;</td>
<td class="xl28">CHINANET-BACKBONE No.31,Jin-rong Street</td>
<td class="xl28" align="right">970</td>
<td class="xl28" align="right">705</td>
<td class="xl29" align="right">-27%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">33182 &nbsp;</td>
<td class="xl28">DIMENOC - HostDime.com, Inc.</td>
<td class="xl28" align="right">940</td>
<td class="xl28" align="right">638</td>
<td class="xl29" align="right">-32%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">7018 &nbsp; &nbsp;</td>
<td class="xl28">ATT-INTERNET4 - AT&amp;T Services, Inc.</td>
<td class="xl28" align="right">934</td>
<td class="xl28" align="right">624</td>
<td class="xl29" align="right">-33%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">24940 &nbsp;</td>
<td class="xl28">HETZNER-AS Hetzner Online AG RZ</td>
<td class="xl28" align="right">872</td>
<td class="xl28" align="right">593</td>
<td class="xl29" align="right">-32%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">26496 &nbsp;</td>
<td class="xl28">AS-26496-GO-DADDY-COM-LLC - GoDaddy.com, LLC</td>
<td class="xl28" align="right">855</td>
<td class="xl28" align="right">560</td>
<td class="xl29" align="right">-35%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">20773 &nbsp;</td>
<td class="xl28">HOSTEUROPE-AS Host Europe GmbH</td>
<td class="xl28" align="right">835</td>
<td class="xl28" align="right">517</td>
<td class="xl29" align="right">-38%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">16276</td>
<td class="xl28">OVH OVH Systems</td>
<td class="xl28" align="right">803</td>
<td class="xl28" align="right">511</td>
<td class="xl29" align="right">-36%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">13768 &nbsp;</td>
<td class="xl28">PEER1 - Peer 1 Network Inc.</td>
<td class="xl28" align="right">707</td>
<td class="xl28" align="right">421</td>
<td class="xl29" align="right">-40%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">14383 &nbsp;</td>
<td class="xl28">VCS-AS - Virtacore Systems Inc</td>
<td class="xl28" align="right">596</td>
<td class="xl28" align="right">420</td>
<td class="xl29" align="right">-30%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">32613 &nbsp;</td>
<td class="xl28">IWEB-AS - iWeb Technologies Inc.</td>
<td class="xl28" align="right">585</td>
<td class="xl28" align="right">367</td>
<td class="xl29" align="right">-37%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">23352 &nbsp;</td>
<td class="xl28">SERVERCENTRAL - Server Central Network</td>
<td class="xl28" align="right">577</td>
<td class="xl28" align="right">350</td>
<td class="xl29" align="right">-39%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">2514 &nbsp; &nbsp;</td>
<td class="xl28">INFOSPHERE NTT PC Communications, Inc.</td>
<td class="xl28" align="right">561</td>
<td class="xl28" align="right">341</td>
<td class="xl29" align="right">-39%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">2519 &nbsp; &nbsp;</td>
<td class="xl28">VECTANT VECTANT Ltd.</td>
<td class="xl28" align="right">531</td>
<td class="xl28" align="right">326</td>
<td class="xl29" align="right">-39%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">15003 &nbsp;</td>
<td class="xl28">NOBIS-TECH - Nobis Technology Group, LLC</td>
<td class="xl28" align="right">521</td>
<td class="xl28" align="right">322</td>
<td class="xl29" align="right">-38%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">22773 &nbsp;</td>
<td class="xl28">ASN-CXA-ALL-CCI-22773-RDC - Cox Communications Inc</td>
<td class="xl28" align="right">484</td>
<td class="xl28" align="right">315</td>
<td class="xl29" align="right">-35%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">6830 &nbsp; &nbsp;</td>
<td class="xl28">LGI-UPC UPC Broadband Holding B.V.</td>
<td class="xl28" align="right">453</td>
<td class="xl28" align="right">307</td>
<td class="xl29" align="right">-32%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">12322 &nbsp;</td>
<td class="xl28">PROXAD Free SAS</td>
<td class="xl28" align="right">449</td>
<td class="xl28" align="right">299</td>
<td class="xl29" align="right">-33%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">21788 &nbsp;</td>
<td class="xl28">NOC - Network Operations Center Inc.</td>
<td class="xl28" align="right">442</td>
<td class="xl28" align="right">295</td>
<td class="xl29" align="right">-33%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">17506 &nbsp;</td>
<td class="xl28">UCOM UCOM Corp.</td>
<td class="xl28" align="right">422</td>
<td class="xl28" align="right">293</td>
<td class="xl29" align="right">-31%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">6939 &nbsp; &nbsp;</td>
<td class="xl28">HURRICANE - Hurricane Electric, Inc.</td>
<td class="xl28" align="right">414</td>
<td class="xl28" align="right">284</td>
<td class="xl29" align="right">-31%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">16265</td>
<td class="xl28">LEASEWEB LeaseWeb B.V.</td>
<td class="xl28" align="right">407</td>
<td class="xl28" align="right">284</td>
<td class="xl29" align="right">-30%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">3269 &nbsp; &nbsp;</td>
<td class="xl28">ASN-IBSNAZ Telecom Italia S.p.a.</td>
<td class="xl28" align="right">402</td>
<td class="xl28" align="right">281</td>
<td class="xl29" align="right">-30%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">29550</td>
<td class="xl28">SIMPLYTRANSIT Simply Transit Ltd</td>
<td class="xl28" align="right">392</td>
<td class="xl28" align="right">271</td>
<td class="xl29" align="right">-31%</td>
</tr>
<tr height="20">
<td class="xl27" height="15">19262 &nbsp;</td>
<td class="xl28">VZGNI-TRANSIT - Verizon Online LLC</td>
<td class="xl28" align="right">390</td>
<td class="xl28" align="right">262</td>
<td class="xl29" align="right">-33%</td>
</tr>

</table>
<p><strong>Kudos</strong></p>
<p>A few other organizations deserve a special shout out for helping with this effort. The great folks at <a href="http://teamcymru.com/" target="_blank">Team Cymru</a> have been tracking open resolvers and other badness online since before CloudFlare was even an idea. Their consistent efforts in this area have been awesome and we're in the process of partnering with them to help get the word out.</p>
<p><span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px;">In addition, SoftLayer has been especially vocal and active in spearheading clean up efforts on its network.&nbsp;</span>As they&nbsp;<a href="http://blog.softlayer.com/2012/the-trouble-with-open-dns-resolvers/" target="_blank">pointed out in a great blog post</a>,<span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px;">&nbsp;because of the size and nature of their network, it's often difficult for them to police the configuration of software their customers run. Even so, they are actively reaching out to customers to educate them about the dangers of running open resolvers on their networks.</span></p>
<p><span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px;">We greatly appreciate country CERTs/CSIRTs and various Information Sharing and Analysis Centers (ISACs) reaching out to us offering to get in touch with some of the less responsive network providers.</span></p>
<p><span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px;">Going forward, we are happy to provide the IP addresses running open resolvers directly to any network provider that is interested in cleaning up their networks. If you're running a network on the list above, please don't hesitate to reach out to us and we'll get you the data you need to help with cleanup.</span></p>
	
</p>

<p><a href="http://blog.cloudflare.com/good-news-open-dns-resolvers-are-getting-clos">Permalink</a> 

	| <a href="http://blog.cloudflare.com/good-news-open-dns-resolvers-are-getting-clos#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/jpeg" height="189" width="267" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-22/acDJcnqgEakaqABIJkcbuHuAAmxqerjoauDvluIjIGcEqunFgmaytJFCxikf/good_news.jpeg">
        <media:thumbnail height="189" width="267" url="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-22/acDJcnqgEakaqABIJkcbuHuAAmxqerjoauDvluIjIGcEqunFgmaytJFCxikf/good_news.jpeg.scaled500.jpg"/>
      </media:content>
    </item>
    <item>
      <pubDate>Wed, 13 Feb 2013 17:29:00 -0800</pubDate>
      <title>When the Bad Guys Name Malware After You, You Know You're Doing Something Right</title>
      <link>http://blog.cloudflare.com/when-the-bad-guys-name-malware-after-you-you</link>
      <guid>http://blog.cloudflare.com/when-the-bad-guys-name-malware-after-you-you</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<a href="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-12/dvvGGAcoACFnrlJmphmdkfprBJqqHtmIDonDcrEhkahGbjBuhijjoFdrzGzm/im_under_attack_page.png.scaled1000.png"><img alt="Im_under_attack_page" height="207" src="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-02-12/dvvGGAcoACFnrlJmphmdkfprBJqqHtmIDonDcrEhkahGbjBuhijjoFdrzGzm/im_under_attack_page.png.scaled500.png" width="500" /></a>
</div>
</p>
<p>CloudFlare's I'm Under Attack Mode (IUAM) is elegantly simple. When a site is under an application layer (Layer 7) distributed denial of service (DDoS) attack, the mode will return a challenge page to a visitor. The challenge requires the visitor's browser to answer a math problem which takes a bit of time to compute. Once successfully answered, the browser can request a cookie and won't be challenged again.</p>
<p><strong>2 + 2 = Surprisingly Effective</strong></p>
<p>IUAM has been incredibly successful at stopping Layer 7 attacks, but it's had a dirty little secret since it was first launched. While we'd suggested that the math problem the browser had to solve would be computationally complex, in reality it was incredibly simple: literally adding together two single-digit integers.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Hard_math" height="215" src="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-12/wdlelijwfbDlHguHDglnmzosbJrbsyxokpqxurBDgmepHJbEHkdJaddhFpJH/hard_math.png.scaled500.png" width="495" />
</div>
</p>
<p>Several people over the last 6 months had written to us to let us know about this "critical vulnerability." They explained how easy it would be for an attacker to reverse engineer the math problem and create malware that could bypass the protection. Internally, we had a bet on how long it would take for some bad guy to actually do so. My money was on "never."</p>
<p><strong>Good News/Bad News: I Lost the Bet</strong></p>
<p>When Lee and I created Project Honey Pot back in 2004 we spent hundreds of engineering hours designing traps that were so random they were hard to identify. Even then, I secretly worried that an enterprising bad guy would recognize some pattern in the traps and be able to avoid them. We watched carefully for 9 years and no one ever took the time to do so. It was great, on one hand, since it meant that Project Honey Pot kept tracking bad guys but, on the other, it meant it was never causing them enough trouble that they'd spend the engineering effort to defeat us. Lee and I learned the lesson: don't over-engineer too early.</p>
<p>Which brings me back to IUAM. This morning we got word from the great folks over at&nbsp;<a href="http://www.eset.com" target="_blank">ESET</a>&nbsp;that they'd <a href="http://www.welivesecurity.com/2013/02/13/malware-evolving-to-defeat-anti-ddos-services-like-cloudflare/" target="_blank">detected malware specifically designed to bypass CloudFlare's IUAM</a>. Called OutFlare -- how cool is it that we have malware named after us!! -- the malware reads our IUAM page, finds the simple math problem, and calculates the answer. It is hardly rocket science, but it was actually pretty thrilling to the whole CloudFlare team that we'd been so successful at stopping bad guys that at least one of them took the time to reverse engineer this protection.</p>
<p><strong>Proof of Work</strong></p>
<p>Unlike me, some other engineers on CloudFlare's team had a suspicion that this day would come. They therefore had, waiting in the wings, code to increase the complexity of IUAM's challenges. The malware pulls the math equation off the page and computes the answer before posting back. The solution was easy: obfuscate the equation and run through some other tricks that make it hard to find the answer if you're not actually rendering the Javascript.</p>
<p>Today, after getting word that the simple version of IUAM had been reverse engineered by the OutFlare's malware, we pushed an update. If you're using IUAM there's nothing you need to do to take advantage of the new protection, we've already updated the protection rendering the OutFlare malware obsolete.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Proof_of_work" height="210" src="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-12/IbamnDezzquIDEBIiblxdHrrkxnlurAtIuwJHplmdrivEJpjrrfceeEbDazc/proof_of_work.jpg.scaled500.jpg" width="240" />
</div>
</p>
<p>Going forward, we have plans if this scheme gets cracked. Specifically, we have an IUAM version that relies on a field of mathematics known as "proof of work" problems. These are difficult to compute answers for but easy to verify. A recent example of such a proof of work problem which has captured the imagination of much of the tech community is Bitcoin. The electronic currency requires a significant amount of computational time to find the answer to a problem, but once found each answer ("coin") is easy to verify.</p>
<p>In Bitcoin's case, the difficulty of the question is adjusted upward over time to compensate for increasing computing power and to control currency inflation. We can use the same premise to increase the "work" that an attacker needs to do when we detect a Layer 7 attack against a CloudFlare customer.</p>
<p><strong>Arms Race? Bet on the Cloud</strong></p>
<p>In these situations there's always a question of whether there will be an arms race between the bad guys writing the malware and the good guys offering protection. In this case there may be, but I like our odds in such a war. As today's example demonstrated, because CloudFlare is deployed as a service and we can update our systems to adjust to new threats in realtime we have an asymmetrical advantage. Pity the poor malware writer who now has to reverse engineer the new IUAM protection and push a code change to all his bots. If he comes up with something effective, we'll just adapt again &mdash; instantly.</p>
<p><div class='p_embed p_image_embed'>
<img alt="Arms_race" height="335" src="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-02-12/fDvlDmuutHqqcuDHloBxcotrybIJbavFABqchrIaFfDEvqvCscfDuxBpcbrF/arms_race.gif.scaled500.gif" width="499" />
</div>
</p>
<p>The history of such arms races suggests you should bet on the cloud to win. In the spam wars, spammers and anti-spam software makers were locked in an arms race that it looked like neither would win from the mid-90s through the mid-2000s. Then something changed: new services like MXLogic, MessageLabs, CloudMark, and Postini started delivering anti-spam not as software but as a "cloud" service. Not only were these services easier to install and administer than previous anti-spam software or appliances, they could also adjust to spammers in realtime. The result has been that today these services have largely won the spam wars.</p>
<p><strong>One More Thing</strong></p>
<p>One more thing with regard to OutFlare. While the malware was able to read and pass the simple math challenge, that is only one layer of IUAM's protection. On the server side, CloudFlare still tracked all requests and, for devices that created a statistically high number of connections, we automatically imposed rate limits and other mitigation techniques. In other words, even without the fix we made, our customers were protected from the attack.</p>
<p>Thanks again to our friends at&nbsp;<a href="http://www.eset.com" target="_blank">ESET</a>&nbsp;for alerting us to the new OutFlare malware. We'll keep our eyes open to any new variants and, as they inevitably arise, we'll continue to adapt to ensure that all CloudFlare customers are always a step ahead of the web's nastiest threats.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/when-the-bad-guys-name-malware-after-you-you">Permalink</a> 

	| <a href="http://blog.cloudflare.com/when-the-bad-guys-name-malware-after-you-you#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="274" width="661" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-02-12/dvvGGAcoACFnrlJmphmdkfprBJqqHtmIDonDcrEhkahGbjBuhijjoFdrzGzm/im_under_attack_page.png">
        <media:thumbnail height="207" width="500" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-02-12/dvvGGAcoACFnrlJmphmdkfprBJqqHtmIDonDcrEhkahGbjBuhijjoFdrzGzm/im_under_attack_page.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="215" width="495" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-12/wdlelijwfbDlHguHDglnmzosbJrbsyxokpqxurBDgmepHJbEHkdJaddhFpJH/hard_math.png">
        <media:thumbnail height="215" width="495" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-12/wdlelijwfbDlHguHDglnmzosbJrbsyxokpqxurBDgmepHJbEHkdJaddhFpJH/hard_math.png.scaled500.png"/>
      </media:content>
      <media:content type="image/jpeg" height="210" width="240" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-12/IbamnDezzquIDEBIiblxdHrrkxnlurAtIuwJHplmdrivEJpjrrfceeEbDazc/proof_of_work.jpg">
        <media:thumbnail height="210" width="240" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-12/IbamnDezzquIDEBIiblxdHrrkxnlurAtIuwJHplmdrivEJpjrrfceeEbDazc/proof_of_work.jpg.scaled500.jpg"/>
      </media:content>
      <media:content type="image/gif" height="335" width="499" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-02-12/fDvlDmuutHqqcuDHloBxcotrybIJbavFABqchrIaFfDEvqvCscfDuxBpcbrF/arms_race.gif">
        <media:thumbnail height="335" width="499" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-02-12/fDvlDmuutHqqcuDHloBxcotrybIJbavFABqchrIaFfDEvqvCscfDuxBpcbrF/arms_race.gif.scaled500.gif"/>
      </media:content>
    </item>
    <item>
      <pubDate>Thu, 07 Feb 2013 21:08:44 -0800</pubDate>
      <title>Facebook Bug Redirects the Web Through Javascript Widget Error</title>
      <link>http://blog.cloudflare.com/facebook-bug-takes-down-much-of-the-web-cloud</link>
      <guid>http://blog.cloudflare.com/facebook-bug-takes-down-much-of-the-web-cloud</guid>
      <description>
        <![CDATA[<p>
	<p>You may have heard that <a href="http://allthingsd.com/20130207/in-one-fell-swoop-apparent-facebook-glitch-deep-sixes-the-web/" target="_blank">Facebook took down a significant portion of the Internet today</a>. A bug in their Facebook Connect script -- which is installed widely across many sites including CNN, MSNBC.com, New York Magazine, and many more places -- caused users to be redirected to a Facebook error page. Here's a video of what it looked like if you visited NBCNews.com:</p>
<p><iframe src="http://www.youtube.com/embed/lcAmokHHuO0" frameborder="0" height="315" width="560"></iframe></p>
<p>The incident raises two good points: 1) the risk of Javascript widgets creating a "single points of failure" on your web page; and 2) the ways in which CloudFlare can help protect you from similar errors.</p>
<p><strong>Widgets &amp; SPOF</strong></p>
<p>Facebook Connect works as a piece of Javascript that is embeded on pages. When the bug occurred, the Javascript effectively hijacked the page and directed it somewhere else. It may seem like installing a widget such as the Facebook button is harmless, but today's incident shows how much harm it can actually cause.</p>
<p><a href="http://twitter.com/BjoernKaiser" target="_blank">Bj&ouml;rn Kaiser</a> wrote a <a href="http://calendar.perfplanet.com/2012/spof-bug/" target="_blank">great blog post last year about the risks that embedded Javascript widgets can create</a>, and how their failure can create a single point of failure (SPOF) on your site. In the post, he describes how you can test the embedded widgets on your page to see what would happen if any of them fail. Given that no widget provider, even Facebook, is infallible it is important to understand the risk of widget failure bringing down your site.</p>
<p><strong>How CloudFlare Helps</strong></p>
<p>There are two distinct ways in which CloudFlare helps protect you from Javascript widgets taking down your site. The first is via our Rocket Loader feature.</p>
<p>While we don't describe it this way often, <a href="http://blog.cloudflare.com/56590463" target="_blank">Rocket Loader</a> is effectively an on-page Javascript optimizer. It sits in front of widgets and makes sure they load as fast as possible. It also has a number of failsafes that can protect from any widget hijacking your site the way Facebook's Connect service did today. While we primarily describe Rocket Loader as a performance feature, in this role it's also very helpful for security and site availability.</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-07/ByGiEzGykCobvvzEabuufAEoBFFxGueajqEqnobAnvrliIJdJwuldzJflakf/rocket.png.scaled1000.png"><img alt="Rocket" height="500" src="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-07/ByGiEzGykCobvvzEabuufAEoBFFxGueajqEqnobAnvrliIJdJwuldzJflakf/rocket.png.scaled500.png" width="500" /></a>
</div>
</p>
<p>The second way we protect sites from misbehaving Javascript widgets is through CloudFlare's app store. Many CloudFlare apps are Javascript widgets of one kind or another. When you install any CloudFlare app, we go through the process of making sure that the app performs well and can run asychronously. This greatly reduces the risk of an CloudFlare-installed app becoming a SPOF. Moreover, because we can install, upgrade, and remove apps centrally, if a problem like Facebook's had occurred with one of the CloudFlare apps, we could quickly remove it from pages to keep it from causing harm.</p>
<p><strong>#savetheweb</strong></p>
<p>Today's Facebook incident shows the risks of misbehaving Javascript widgets, but it also helps drive home the point on how CloudFlare is really building a better web. To that end, we will continue to invest in improving Rocket Loader and adding more and more apps to the CloudFlare Apps Marketplace. If you haven't turned on Rocket Loader or added an app through the <a href="http://www.cloudflare.com/apps" target="_blank">CloudFlare Apps Marketplace</a>, you now have one more reason check them both out.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/facebook-bug-takes-down-much-of-the-web-cloud">Permalink</a> 

	| <a href="http://blog.cloudflare.com/facebook-bug-takes-down-much-of-the-web-cloud#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="792" width="792" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-02-07/ByGiEzGykCobvvzEabuufAEoBFFxGueajqEqnobAnvrliIJdJwuldzJflakf/rocket.png">
        <media:thumbnail height="500" width="500" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-07/ByGiEzGykCobvvzEabuufAEoBFFxGueajqEqnobAnvrliIJdJwuldzJflakf/rocket.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Mon, 04 Feb 2013 06:26:00 -0800</pubDate>
      <title>New &quot;Lucky Thirteen&quot; SSL Vulnerabilities: CloudFlare Users Protected</title>
      <link>http://blog.cloudflare.com/new-ssl-vulnerabilities-cloudflare-users-prot</link>
      <guid>http://blog.cloudflare.com/new-ssl-vulnerabilities-cloudflare-users-prot</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<img alt="Cloudflare_secure_ssl" height="380" src="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-03/zfwigkxGvowpnsCrEgIIelpsvigbleHJnqkImviCpHGguaBHabGcgdFyitFx/cloudflare_secure_ssl.png.scaled500.png" width="311" />
</div>
</p>
<p>CloudFlare often gets early word of new vulnerabilities before they are released. Last week we got word that today (Monday, February 4, 2013) there would be a new SSL vulnerability announced. This vulnerability follows the BEAST and CRIME vulnerabilities that have been discovered over the last 18 months. The bad news is that TLS 1.1/1.2 do not fix the issue.</p>
<p>The vulnerabilities are known as the <a href="http://www.isg.rhul.ac.uk/tls/">Lucky Thirteen</a>.</p>
<p><div class='p_embed p_image_embed'>
<img alt="California_13" height="401" src="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-02-04/fJIChqltqzkjfkuBirbrpznInrnxeukuBiytpAfuldjoheAmtnaqbcgzmsdc/California_13.png.scaled500.png" width="385" />
</div>
</p>
<p>The good news is that our analysis of the newest vulnerability suggests that, while theoretically possible, it is fairly difficult to exploit. It is a timing attack and you'd need to create a fairly large number of connections and measure the differences in timing. That's possible, but non-trivial.</p>
<p>That said, at CloudFlare we want to ensure that even remote risks are fully mitigated. In this case, the good news is CloudFlare's SSL configuration is, by default, not generally vulnerable to the new attack. Specifically, because we deprioritize the vulnerable SSL cipher, it makes anyone using a modern browser invulnerable to the attack when visiting a CloudFlare-protected site over an SSL connection.</p>
<p>While the easiest way to ensure that your site is protected from the new vulnerability is to sign up for CloudFlare's service, if you haven't gotten around to that yet then there are some steps you should take. First, when a new version of OpenSSL is released that removes this vulnerability, which we expect will happen in the next few weeks, you should upgrade. Second, you should prioritize the RC4 cipher in your web server above others as it isn't vulnerable.</p>
<p>Here's the Apache SSL cipher suite configuration we'd recommend:</p>
<div class="CodeRay">
  <div class="code"><pre>SSLProtocol -ALL +SSLv3 +TLSv1
SSLCipherSuite ECDHE-RSA-AES128-SHA256:AES128-GCM-SHA256:RC4:HIGH:!MD5:!aNULL:!EDH
SSLHonorCipherOrder on</pre></div>
</div>

<p>Here's the NGINX SSL cyber suite configuration we'd recommend:&nbsp;</p>
<div class="CodeRay">
  <div class="code"><pre>ssl_protocols               SSLv3 TLSv1 TLSv1.1 TLSv1.2;
ssl_ciphers                 ECDHE-RSA-AES128-SHA256:AES128-GCM-SHA256:RC4:HIGH:!MD5:!aNULL:!EDH;
ssl_prefer_server_ciphers   on;</pre></div>
</div>
	
</p>

<p><a href="http://blog.cloudflare.com/new-ssl-vulnerabilities-cloudflare-users-prot">Permalink</a> 

	| <a href="http://blog.cloudflare.com/new-ssl-vulnerabilities-cloudflare-users-prot#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725058/15-00-35-med2.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbOrGDRWV</posterous:profileUrl>
        <posterous:firstName>Matthew</posterous:firstName>
        <posterous:lastName>Prince</posterous:lastName>
        <posterous:nickName>eastdakota</posterous:nickName>
        <posterous:displayName>Matthew Prince</posterous:displayName>
      </posterous:author>
      <media:content type="image/png" height="109" width="566" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-03/kiukihaeksxqrnICslrjGpJjryxuDxFBBHgvxAIkEqfqalEAItfvyJuiGFqJ/rc4_ssl_message.png">
        <media:thumbnail height="96" width="500" url="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-02-03/kiukihaeksxqrnICslrjGpJjryxuDxFBBHgvxAIkEqfqalEAItfvyJuiGFqJ/rc4_ssl_message.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="380" width="311" url="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-02-03/zfwigkxGvowpnsCrEgIIelpsvigbleHJnqkImviCpHGguaBHabGcgdFyitFx/cloudflare_secure_ssl.png">
        <media:thumbnail height="380" width="311" url="http://getfile0.posterous.com/getfile/files.posterous.com/temp-2013-02-03/zfwigkxGvowpnsCrEgIIelpsvigbleHJnqkImviCpHGguaBHabGcgdFyitFx/cloudflare_secure_ssl.png.scaled500.png"/>
      </media:content>
      <media:content type="image/png" height="401" width="385" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-02-04/fJIChqltqzkjfkuBirbrpznInrnxeukuBiytpAfuldjoheAmtnaqbcgzmsdc/California_13.png">
        <media:thumbnail height="401" width="385" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-02-04/fJIChqltqzkjfkuBirbrpznInrnxeukuBiytpAfuldjoheAmtnaqbcgzmsdc/California_13.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Fri, 01 Feb 2013 14:27:00 -0800</pubDate>
      <title>Edge Cache Expire TTL: Easiest way to override any existing headers</title>
      <link>http://blog.cloudflare.com/edge-cache-expire-ttl-easiest-way-to-override</link>
      <guid>http://blog.cloudflare.com/edge-cache-expire-ttl-easiest-way-to-override</guid>
      <description>
        <![CDATA[<p>
	<p><div class='p_embed p_image_embed'>
<img alt="Cache_rules_everything_around_me" height="320" src="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-01/IrphbphjygcgtHjvvFhuudCsBldkvialAeEinnnbhaksvBhsusxqvamvsngF/Cache_Rules_Everything_Around_Me.png.scaled500.png" width="299" />
</div>
</p>
<p>CloudFlare makes caching easy. Our service automatically determines what files to cache based on file extensions. Performance benefits kick in automatically.</p>
<p>For customers that want advanced caching, beyond the defaults, we have <a href="http://blog.cloudflare.com/introducing-pagerules-advanced-caching" target="_blank">Cache Everything</a>&nbsp;available as Page Rules. Designate a URL and CloudFlare will cache everything, including HTML, out at the edges of our global network.</p>
<p>With Cache Everything, we respect all headers. If there is any header in place from the server or a CMS solution like WordPress, we will respect it. However, we got many requests from customers who wanted an easy way to override any existing headers. Today, we are releasing a new feature called 'Edge cache expire TTL' that does just that.</p>
<div>
<p><strong>What is Edge Cache Expire TTL?</strong></p>
<p><span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px;">Edge cache expire TTL is the setting that controls how long CloudFlare's edge servers will cache a resource before requesting a fresh copy from your server. When you create a Cache Everything Page Rule, you now may choose whether to respect all existing headers or to override any headers that are in place from your server. By overwriting the headers, CloudFlare will cache more content at the CloudFlare edge network, decreasing load to your server.</span><span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px;">&nbsp;</span></p>
<p><span style="color: #222222; font-family: arial, sans-serif;">Common situations where you may choose to overwrite existing headers:</span></p>
<div style="margin: 0px; padding: 0px;">
<div style="line-height: 1.45em; margin: 0px; padding: 0px; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;">
<ul style="margin: 0px; padding: 2px 0px 5px 36px;">
<li style="margin: 5px 11px 5px 0px; padding: 0px; background-image: none;"><span style="margin: 0px; padding: 0px; line-height: 1.45em;">You expect a large surge in traffic</span></li>
<li style="margin: 5px 11px 5px 0px; padding: 0px; background-image: none;"><span style="margin: 0px; padding: 0px; line-height: 1.45em;">You are under DDOS attack</span></li>
<li style="margin: 5px 11px 5px 0px; padding: 0px; background-image: none;"><span style="margin: 0px; padding: 0px; line-height: 1.45em;">You are not sure what the headers on WordPress or your server are set to</span></li>
<li style="margin: 5px 11px 5px 0px; padding: 0px; background-image: none;"><span style="margin: 0px; padding: 0px; line-height: 1.45em;">You are using WordPress and want to easily overwrite the default settings &nbsp;</span></li>
</ul>
</div>
</div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;">It is important to emphasize when you&nbsp;<em>do not</em>&nbsp;want to use Cache Everything. If you have any personalized information on the page like login information or credit card information, you do not want to use the Cache Everything option.</div>
</div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;">
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><strong style="line-height: 1.45em;">What is Browser Cache Expire TTL?</strong></div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><span style="color: #222222; font-family: arial, sans-serif;">Browser cache expire TTL is the time that CloudFlare instructs a visitor's browser to cache a resource. Until this time expires, the browser will load the resource from its local cache thus speeding up the request significantly. CloudFlare will respect the headers that you give us from your web server, and then we will communicate with the browser based on the time selected in this drop down menu.</span></div>
<p />
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><strong style="line-height: 1.45em;">Using both Edge Cache Expire TTL and Browser Cache Expire TTL</strong></div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px; line-height: normal;">When you'd like to have CloudFlare cache your content but want your visitors to always get a fresh copy of the page, you can use the new 'Edge cache expire TTL' setting to express this differentiation. Set a value for 'Edge cache expire TTL' to how often you want the CloudFlare CDN to refresh from your server, and 'Browser cache expire TTL' to how often you want your visitors' browsers to refresh the page content. This is useful when you have a rapidly changing page but still want the benefit of the CloudFlare cache to reduce your server load.</span></div>
<p />
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><span style="color: #222222; font-family: arial, sans-serif; font-size: 13.333333969116211px; line-height: normal;"> </span>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><strong>Plan Details</strong></div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><span style="line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;">CloudFlare offers a range of edge cache expire TTLs based on plan type:</span></div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;">
<div style="margin: 0px; padding: 0px; line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;">
<ul>
<li><span style="line-height: 1.45em;">Free &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;2 hours</span></li>
<li><span style="line-height: 1.45em;">Pro &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;1 hour</span></li>
<li><span style="line-height: 1.45em;">Business &nbsp; &nbsp;30 minutes</span></li>
<li><span style="line-height: 1.45em;">Enterprise &nbsp; as low as 30 seconds&nbsp;</span></li>
</ul>
</div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;">
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><span style="line-height: 1.45em;">A Pro customer may set the refetch time to 1 hour. After 60 minutes, we return to your server for a fresh copy of the resource. Business customers may lower the refetch interval to 30 minutes. Enterprise customers may set this interval as low as 30 seconds.</span></div>
<p />
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><span style="line-height: 1.45em;"> </span>
<div style="margin: 0px; padding: 0px; line-height: 1.45em;">
<div style="margin: 0px; padding: 0px; line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;">
<div style="margin: 0px; padding: 0px; line-height: 1.45em;"><strong style="line-height: 1.45em;">How to Turn It On</strong></div>
</div>
</div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;">Login in to your CloudFlare account and choose "Page Rules" from under the gear icon. Enter the URL that you want to Cache Everything (under Custom Caching):</div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;"><div class='p_embed p_image_embed'>
<a href="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-01/fJblguGlEdplvIvCcJmmkpwgcIuAoBeAEBkiAbywCgFatdfnucAEybGdzzep/Cache_Everything.tiff.scaled1000.jpg"><img alt="Cache_everything" height="142" src="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-02-01/fJblguGlEdplvIvCcJmmkpwgcIuAoBeAEBkiAbywCgFatdfnucAEybGdzzep/Cache_Everything.tiff.scaled500.jpg" width="500" /></a>
</div>
The edge cache server TTL option will appear:</div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;"><div class='p_embed p_image_embed'>
<a href="http://getfile7.posterous.com/getfile/files.posterous.com/temp-2013-02-01/IqjFqpmngecGjEifEgqlpveChegBthGqwFvJFfeksfqflpwCtempcfzgGcvE/Edge_cache_expire_TTL_appears.tiff.scaled1000.jpg"><img alt="Edge_cache_expire_ttl_appears" height="267" src="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-02-01/IqjFqpmngecGjEifEgqlpveChegBthGqwFvJFfeksfqflpwCtempcfzgGcvE/Edge_cache_expire_TTL_appears.tiff.scaled500.jpg" width="500" /></a>
</div>
The default setting is set to "Respect all existing headers." To override this setting, choose a time from the drop down menu:</div>
<div style="margin: 0px; padding: 0px; line-height: 1.45em; color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif; font-size: 12px;"><div class='p_embed p_image_embed'>
<a href="http://getfile4.posterous.com/getfile/files.posterous.com/temp-2013-02-01/gaEHhHdIgfDwlsaCwHsrBqCeEkJhHvqscCtfuraFwbdvDEIEnDcdmmdgFnyx/Edge_cache_expire_TTL_dropdown.tiff.scaled1000.jpg"><img alt="Edge_cache_expire_ttl_dropdown" height="361" src="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-01/gaEHhHdIgfDwlsaCwHsrBqCeEkJhHvqscCtfuraFwbdvDEIEnDcdmmdgFnyx/Edge_cache_expire_TTL_dropdown.tiff.scaled500.jpg" width="500" /></a>
</div>

<p style="color: #000000; font-family: Arial, Helvetica, sans-serif; font-size: 13px; line-height: normal;"><span style="color: #333333; font-family: Helvetica Neue, Helvetica, Arial, sans-serif;"><span style="font-size: 12px; line-height: 17.390625px;">You can find more information in our knowledge base articles&nbsp;<a href="https://support.cloudflare.com/entries/23023893-what-does-edge-cache-expire-ttl-mean" target="_blank">here</a>&nbsp;and&nbsp;<a href="https://support.cloudflare.com/entries/23009261-what-does-browser-cache-expire-ttl-mean" target="_blank">here.</a></span></span></p>
<p style="color: #000000; font-family: Arial, Helvetica, sans-serif; font-size: 13px; line-height: normal;">Give it a try and let us know what you think.</p>
</div>
</div>
<p />
</div>
</div>
</div>
</div>
<p>&nbsp;</p>
	
</p>

<p><a href="http://blog.cloudflare.com/edge-cache-expire-ttl-easiest-way-to-override">Permalink</a> 

	| <a href="http://blog.cloudflare.com/edge-cache-expire-ttl-easiest-way-to-override#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/725124/IMG_0043.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/4xgdbv0Udt1n</posterous:profileUrl>
        <posterous:firstName>Michelle</posterous:firstName>
        <posterous:lastName>Zatlyn</posterous:lastName>
        <posterous:nickName>Michelle Z.</posterous:nickName>
        <posterous:displayName>Michelle Zatlyn</posterous:displayName>
      </posterous:author>
      <media:content type="image/tiff" height="208" width="734" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-02-01/fJblguGlEdplvIvCcJmmkpwgcIuAoBeAEBkiAbywCgFatdfnucAEybGdzzep/Cache_Everything.tiff">
        <media:thumbnail height="142" width="500" url="http://getfile5.posterous.com/getfile/files.posterous.com/temp-2013-02-01/fJblguGlEdplvIvCcJmmkpwgcIuAoBeAEBkiAbywCgFatdfnucAEybGdzzep/Cache_Everything.tiff.scaled500.jpg"/>
      </media:content>
      <media:content type="image/tiff" height="417" width="782" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-01/IqjFqpmngecGjEifEgqlpveChegBthGqwFvJFfeksfqflpwCtempcfzgGcvE/Edge_cache_expire_TTL_appears.tiff">
        <media:thumbnail height="267" width="500" url="http://getfile2.posterous.com/getfile/files.posterous.com/temp-2013-02-01/IqjFqpmngecGjEifEgqlpveChegBthGqwFvJFfeksfqflpwCtempcfzgGcvE/Edge_cache_expire_TTL_appears.tiff.scaled500.jpg"/>
      </media:content>
      <media:content type="image/tiff" height="508" width="704" url="http://getfile9.posterous.com/getfile/files.posterous.com/temp-2013-02-01/gaEHhHdIgfDwlsaCwHsrBqCeEkJhHvqscCtfuraFwbdvDEIEnDcdmmdgFnyx/Edge_cache_expire_TTL_dropdown.tiff">
        <media:thumbnail height="361" width="500" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-01/gaEHhHdIgfDwlsaCwHsrBqCeEkJhHvqscCtfuraFwbdvDEIEnDcdmmdgFnyx/Edge_cache_expire_TTL_dropdown.tiff.scaled500.jpg"/>
      </media:content>
      <media:content type="image/png" height="320" width="299" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-02-01/IrphbphjygcgtHjvvFhuudCsBldkvialAeEinnnbhaksvBhsusxqvamvsngF/Cache_Rules_Everything_Around_Me.png">
        <media:thumbnail height="320" width="299" url="http://getfile6.posterous.com/getfile/files.posterous.com/temp-2013-02-01/IrphbphjygcgtHjvvFhuudCsBldkvialAeEinnnbhaksvBhsusxqvamvsngF/Cache_Rules_Everything_Around_Me.png.scaled500.png"/>
      </media:content>
    </item>
    <item>
      <pubDate>Fri, 18 Jan 2013 01:23:00 -0800</pubDate>
      <title>WordPress London Meetup January 2013</title>
      <link>http://blog.cloudflare.com/wordpress-london-meetup-january-2013</link>
      <guid>http://blog.cloudflare.com/wordpress-london-meetup-january-2013</guid>
      <description>
        <![CDATA[<p>
	<p>Last night I gave a short presentation about how to use CloudFlare with WordPress sites to about 60 people attending the <a href="http://www.meetup.com/London-WordPress/events/81910532/">WordPress London Meetup</a>. CloudFlare was happy to be sponsor of the event providing drinks, beers and lots and lots of pizza. The meetup was held at the <a href="http://www.campuslondon.com">Google Campus</a>.</p>
<p><div class='p_embed p_image_embed'>
<a href="http://getfile8.posterous.com/getfile/files.posterous.com/temp-2013-01-18/qqjdxgjjJbEdGngAiuagdoywswrGoJABcHJmongztFlffaFBywvkbhnEiFrd/IMG_4277.JPG.scaled1000.jpg"><img alt="Img_4277" height="375" src="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-01-18/qqjdxgjjJbEdGngAiuagdoywswrGoJABcHJmongztFlffaFBywvkbhnEiFrd/IMG_4277.JPG.scaled500.jpg" width="500" /></a>
</div>
There were two talks: I was preceded by designer <a href="http://laurakalbag.com/">Laura Kalbag</a> who talked about designing icons for WordPress sites. This is something that she made look incredibly easy using a tool called <a href="http://www.bohemiancoding.com/sketch/">Sketch</a>. I suspect that however good Sketch is, I'd end up drawing icons that looked awful!</p>
<p>My talk was about using WordPress and CloudFlare together. CloudFlare has a ton of features and I highlighted some that are of great interest to WordPress users including the <a href="http://wordpress.org/extend/plugins/cloudflare/">CloudFlare WordPress Plugin</a>&nbsp;and&nbsp;our integration with <a href="http://blog.cloudflare.com/w3-total-cache-w3tc-total-cloudflare-integrat">W3TC</a>.</p>
<p>The other features that people found most interesting were:</p>
<ol>
<li><a href="http://blog.cloudflare.com/always-online-v2">Always Online</a>: CloudFlare crawls the WordPress site and keeps a copy in a special cache. If the original site goes down CloudFlare serves up the most recent version from the crawler cache with a banner indicating that it is old content. This helps keep sites online when things go badly wrong.</li>
<li><a href="http://blog.cloudflare.com/an-all-new-and-improved-autominify">Auto-minify</a>: many WordPress sites have large amount of HTML, CSS and JavaScript (especially if they use lots of plugins). Auto-minify helps shrink those resources so that sites are delivered faster to web browsers.</li>
<li><a href="http://blog.cloudflare.com/56590463">Rocket Loader</a>: a tool that reorganizes the loading of resources such as CSS and JavaScript to that they are downloaded to web browsers quickly by bundling them.</li>
<li>A new, unannounced feature that I'm calling "Help, I've gone viral!" which allows any web site owner to instantly tell CloudFlare to start completely caching a URL (overriding any caching headers set by the site) to cope with load. With this if a URL goes viral and is overloading a WordPress site it's possible to just paste in its URL and ask CloudFlare to take the load of that page. We'll be writing more about that feature when it's released.</li>
</ol>
<p>And, of course, other CloudFlare features like <a href="http://blog.cloudflare.com/easiest-ssl-ever-now-included-automatically-w">Easy SSL</a>, <a href="http://blog.cloudflare.com/spdy-now-one-click-simple-for-any-website">SPDY</a>, and <a href="http://blog.cloudflare.com/introducing-cloudflares-automatic-ipv6-gatewa">IPv6</a> help everyone get the latest technology onto their site quickly.</p>
	
</p>

<p><a href="http://blog.cloudflare.com/wordpress-london-meetup-january-2013">Permalink</a> 

	| <a href="http://blog.cloudflare.com/wordpress-london-meetup-january-2013#comment">Leave a comment&nbsp;&nbsp;&raquo;</a>

</p>]]>
      </description>
      <posterous:author>
        <posterous:userImage>http://files.posterous.com/user_profile_pics/1943163/jgc.jpg</posterous:userImage>
        <posterous:profileUrl>http://posterous.com/users/lAAjkS0LMSs0G</posterous:profileUrl>
        <posterous:firstName>John </posterous:firstName>
        <posterous:lastName>Graham-Cumming</posterous:lastName>
        <posterous:nickName>jgrahamc</posterous:nickName>
        <posterous:displayName>John  Graham-Cumming</posterous:displayName>
      </posterous:author>
      <media:content type="image/jpeg" height="2448" width="3264" url="http://getfile3.posterous.com/getfile/files.posterous.com/temp-2013-01-18/qqjdxgjjJbEdGngAiuagdoywswrGoJABcHJmongztFlffaFBywvkbhnEiFrd/IMG_4277.JPG">
        <media:thumbnail height="375" width="500" url="http://getfile1.posterous.com/getfile/files.posterous.com/temp-2013-01-18/qqjdxgjjJbEdGngAiuagdoywswrGoJABcHJmongztFlffaFBywvkbhnEiFrd/IMG_4277.JPG.scaled500.jpg"/>
      </media:content>
    </item>
  </channel>
</rss>
`

const feed1 = `
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" media="screen" href="/~d/styles/rss2full.xsl"?><?xml-stylesheet type="text/css" media="screen" href="http://feeds.arstechnica.com/~d/styles/itemcontent.css"?><rss xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:wfw="http://wellformedweb.org/CommentAPI/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:sy="http://purl.org/rss/1.0/modules/syndication/" xmlns:slash="http://purl.org/rss/1.0/modules/slash/" xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" version="2.0">

<channel>
	<title>Ars Technica » Scientific Method</title>
	
	<link>http://arstechnica.com</link>
	<description>The Art of Technology</description>
	<lastBuildDate>Sat, 20 Apr 2013 16:01:59 +0000</lastBuildDate>
	<language>en-US</language>
	<sy:updatePeriod>hourly</sy:updatePeriod>
	<sy:updateFrequency>1</sy:updateFrequency>
	<generator>http://wordpress.org/?v=3.5.1</generator>
		<atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="self" type="application/rss+xml" href="http://feeds.arstechnica.com/arstechnica/science" /><feedburner:info uri="arstechnica/science" /><atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="hub" href="http://pubsubhubbub.appspot.com/" /><feedburner:feedFlare href="http://www.plusmo.com/add?url=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://plusmo.com/res/graphics/fbplusmo.gif">Subscribe with Plusmo</feedburner:feedFlare><feedburner:feedFlare href="http://www.thefreedictionary.com/_/hp/AddRSS.aspx?http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://img.tfd.com/hp/addToTheFreeDictionary.gif">Subscribe with The Free Dictionary</feedburner:feedFlare><feedburner:feedFlare href="http://www.bitty.com/manual/?contenttype=rssfeed&amp;contentvalue=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.bitty.com/img/bittychicklet_91x17.gif">Subscribe with Bitty Browser</feedburner:feedFlare><feedburner:feedFlare href="http://www.newsalloy.com/?rss=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.newsalloy.com/subrss3.gif">Subscribe with NewsAlloy</feedburner:feedFlare><feedburner:feedFlare href="http://www.live.com/?add=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://tkfiles.storage.msn.com/x1piYkpqHC_35nIp1gLE68-wvzLZO8iXl_JMledmJQXP-XTBOLfmQv4zhj4MhcWEJh_GtoBIiAl1Mjh-ndp9k47If7hTaFno0mxW9_i3p_5qQw">Subscribe with Live.com</feedburner:feedFlare><feedburner:feedFlare href="http://mix.excite.eu/add?feedurl=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://image.excite.co.uk/mix/addtomix.gif">Subscribe with Excite MIX</feedburner:feedFlare><feedburner:feedFlare href="http://download.attensa.com/app/get_attensa.html?feedurl=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.attensa.com/blogs/attensa/WindowsLiveWriter/BadgeredintoBadges_10C02/attensa_feed_button5.gif">Subscribe with Attensa for Outlook</feedburner:feedFlare><feedburner:feedFlare href="http://www.webwag.com/wwgthis.php?url=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.webwag.com/images/wwgthis.gif">Subscribe with Webwag</feedburner:feedFlare><feedburner:feedFlare href="http://www.podcastready.com/oneclick_bookmark.php?url=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.podcastready.com/images/podcastready_button.gif">Subscribe with Podcast Ready</feedburner:feedFlare><feedburner:feedFlare href="http://www.flurry.com/pushRssFeed.do?r=fb&amp;url=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.flurry.com/images/flurry_rss_logo2.gif">Subscribe with Flurry</feedburner:feedFlare><feedburner:feedFlare href="http://www.wikio.com/subscribe?url=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.wikio.com/shared/img/add2wikio.gif">Subscribe with Wikio</feedburner:feedFlare><feedburner:feedFlare href="http://www.dailyrotation.com/index.php?feed=http%3A%2F%2Ffeeds.arstechnica.com%2Farstechnica%2Fscience" src="http://www.dailyrotation.com/rss-dr2.gif">Subscribe with Daily Rotation</feedburner:feedFlare><feedburner:browserFriendly>This is an XML content feed. It is intended to be viewed in a newsreader or syndicated to another site, subject to copyright and fair use.</feedburner:browserFriendly><item>
		<title>Interactive map plots locations of more than 100 million species</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/5U2zfVNz7U0/</link>
		<comments>http://arstechnica.com/science/2013/04/interactive-map-plots-locations-of-more-than-100-million-species/#comments</comments>
		<pubDate>Sat, 20 Apr 2013 14:05:35 +0000</pubDate>
		<dc:creator>WIRED UK</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[biodiversity]]></category>
		<category><![CDATA[endangered species]]></category>
		<category><![CDATA[USGS]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=249715</guid>
		<description><![CDATA[It's the most comprehensive map of US biodiversity ever made.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>The <a href="http://www.usgs.gov/">United States Geological Society</a> (USGS) has launched an online database and map that keeps track of more than 100 million different species and where they live within the United States,</p>
<p>Biodiversity Serving Our Nation, or <a href="http://bison.usgs.ornl.gov/">BISON</a> (a backronym if ever there was one), contains location-specific records of where living species are within the US. Its data comes from hundreds of different organisations and thousands of scientists, making it the most comprehensive map of American biodiversity ever made.</p>
<p>Anyone can search by scientific or common name of any living species (plant or animal), and can look to see what lives within any specific geographic area they want by drawing a perimeter—so, for example, searching to see exactly which forests in Virginia have been infected with a tree fungus.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/interactive-map-plots-locations-of-more-than-100-million-species/#p3">Read 6 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/interactive-map-plots-locations-of-more-than-100-million-species/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=5U2zfVNz7U0:Sxy21Q_2poE:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=5U2zfVNz7U0:Sxy21Q_2poE:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=5U2zfVNz7U0:Sxy21Q_2poE:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=5U2zfVNz7U0:Sxy21Q_2poE:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=5U2zfVNz7U0:Sxy21Q_2poE:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=5U2zfVNz7U0:Sxy21Q_2poE:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/5U2zfVNz7U0" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/interactive-map-plots-locations-of-more-than-100-million-species/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/interactive-map-plots-locations-of-more-than-100-million-species/</feedburner:origLink></item>
		<item>
		<title>This new microbattery is interesting, but not as good as the hype</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/_C3kkuXeycQ/</link>
		<comments>http://arstechnica.com/science/2013/04/this-new-microbattery-is-interesting-but-not-as-good-as-the-hype/#comments</comments>
		<pubDate>Fri, 19 Apr 2013 17:55:25 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[batteries]]></category>
		<category><![CDATA[materials science]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=249607</guid>
		<description><![CDATA[Tiny battery delivers power like a supercapacitor, but performance fades fast.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/Screen-Shot-2013-04-19-at-1.10.14-PM-640x401.png"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">The alternating metal foam electrodes. Scale bar is 50 micrometers.</div>
	
			<div class="caption-credit">
							From Pikul et. al., referenced here.				</div>
	  </div>
  </div>
 <p>Battery technology is advancing fast—just compare the Watt-hours you can stuff into current devices to what was on the market a few decades ago (the Powerbook 140 had a 4200 mAh battery, while the current iPad has nearly three times that capacity). Research is continuing to advance at a rapid pace, which is why it sometimes seems like we cover a new miracle battery every month. We try to make it clear that most of these are still years away from commercialization, and some technologies will just never get there.</p>
<p>It's with that caution in mind that we turn to a battery development that was even picked up <a href="http://www.bbc.co.uk/news/technology-22191650">by the BBC</a> earlier this week. In a new paper, researchers have reported that they are developing a microscopic battery that's suitable for integrating directly into electronics, just as capacitors now are. Compared to existing capacitors, its performance is impressive. That announcement doesn't mean, however, that these things can be scaled up to work as full-sized batteries that can power gadgets or cars. There are also a few hitches, like rapidly degrading performance, that need to be worked out.</p>
<p>That said, making the batteries is a pretty impressive process. In general, the batteries we buy are made from bulk processes, where materials self-organize into the different parts necessary to get the battery to work. This self-organization can be directed to a greater or lesser degree, but it's still dependent upon the material properties of the battery components.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/this-new-microbattery-is-interesting-but-not-as-good-as-the-hype/#p3">Read 8 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/this-new-microbattery-is-interesting-but-not-as-good-as-the-hype/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=_C3kkuXeycQ:VDAnom6a93I:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=_C3kkuXeycQ:VDAnom6a93I:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=_C3kkuXeycQ:VDAnom6a93I:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=_C3kkuXeycQ:VDAnom6a93I:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=_C3kkuXeycQ:VDAnom6a93I:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=_C3kkuXeycQ:VDAnom6a93I:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/_C3kkuXeycQ" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/this-new-microbattery-is-interesting-but-not-as-good-as-the-hype/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/this-new-microbattery-is-interesting-but-not-as-good-as-the-hype/</feedburner:origLink></item>
		<item>
		<title>Barcodes let scientists track every ant in a colony</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/J2a3u7h4OXE/</link>
		<comments>http://arstechnica.com/science/2013/04/barcodes-let-scientists-track-every-ant-in-a-colony/#comments</comments>
		<pubDate>Fri, 19 Apr 2013 16:00:00 +0000</pubDate>
		<dc:creator>Ars Staff</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[ants]]></category>
		<category><![CDATA[Biology]]></category>
		<category><![CDATA[entomology]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=249343</guid>
		<description><![CDATA[Study tracks workers as they change jobs with age.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/mersch1HR-640x719.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">A barcoded ant gets to work.</div>
	
			<div class="caption-credit">
							Alessandro Crespi				</div>
	  </div>
  </div>
 <p>For creatures with very small brains, ants build strikingly complex societies. How a colony of hundreds or thousands of ants maintains order remains poorly understood, but new high tech research methods might be able to shed some light on the complexity of the colony.</p>
<p>A team of Swiss scientists glued barcodes to hundreds of ants living in six laboratory colonies and recorded all of their movements for more than a month. The video analysis recorded the position and orientation of every ant, twice a second. They published the results of the ant tracking in <em>Science</em> this week.</p>
<p>What can you learn from watching 9 million ant-to-ant interactions?</p>
</div><p><a href="http://arstechnica.com/science/2013/04/barcodes-let-scientists-track-every-ant-in-a-colony/#p3">Read 8 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/barcodes-let-scientists-track-every-ant-in-a-colony/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=J2a3u7h4OXE:-aIsBXBRsA8:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=J2a3u7h4OXE:-aIsBXBRsA8:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=J2a3u7h4OXE:-aIsBXBRsA8:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=J2a3u7h4OXE:-aIsBXBRsA8:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=J2a3u7h4OXE:-aIsBXBRsA8:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=J2a3u7h4OXE:-aIsBXBRsA8:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/J2a3u7h4OXE" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/barcodes-let-scientists-track-every-ant-in-a-colony/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/barcodes-let-scientists-track-every-ant-in-a-colony/</feedburner:origLink></item>
		<item>
		<title>Poor, cute bunnies likely to get eaten when the snow melts early</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/Vf1FrLWTFeY/</link>
		<comments>http://arstechnica.com/science/2013/04/poor-cute-bunnies-likely-to-get-eaten-when-the-snow-melts-early/#comments</comments>
		<pubDate>Fri, 19 Apr 2013 14:00:34 +0000</pubDate>
		<dc:creator>Diana Gitig</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Biology]]></category>
		<category><![CDATA[climate change]]></category>
		<category><![CDATA[Earth sciences]]></category>
		<category><![CDATA[Ecology]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247011</guid>
		<description><![CDATA[Hares change coat color for winter based on the calendar, not the conditions.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/82C795C8-1DD8-B71C-07793A6439519ECE-640x480.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.nps.gov/common/uploads/photogallery/akr/park/gaar/82C795C8-1DD8-B71C-07793A6439519ECE/82C795C8-1DD8-B71C-07793A6439519ECE.jpg">National Park Service</a>
				</div>
	  </div>
  </div>
 <p><em>Animals that turn white in winter are having a mismatch with the reduced snowpack in their environments.</em></p>
<p>Measuring the impact of climate change on animals is difficult, because (a) climate change is complex and (b) animals are complex. Climate change can alter the environment in many different ways, and each of these changes can affect the food supplies, hibernation patterns, reproductive behavior, and migratory patterns of different animal species. The changes in the animals all affect each other, too, since many of them are interrelated in food chains and webs that can be hard to disentangle.</p>
<p>Luckily, researchers at the University of Montana stumbled upon one simple, obvious part of an animal that they could measure in response to a simple, obvious change in climate. Snowshoe hares change color from brown to white in the winter so they can be camouflaged in the snow. So do at least nine other species of cute, cuddly mammals: Arctic foxes, collared lemmings, long-tailed weasels, stoats, mountain hares, Arctic hares, white-tailed jackrabbits, Siberian hamsters, and least weasels. (Bet you didn't know that hamsters are tough enough to handle winter in Siberia.)</p>
</div><p><a href="http://arstechnica.com/science/2013/04/poor-cute-bunnies-likely-to-get-eaten-when-the-snow-melts-early/#p3">Read 5 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/poor-cute-bunnies-likely-to-get-eaten-when-the-snow-melts-early/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Vf1FrLWTFeY:R2K50ohef9Q:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=Vf1FrLWTFeY:R2K50ohef9Q:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Vf1FrLWTFeY:R2K50ohef9Q:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=Vf1FrLWTFeY:R2K50ohef9Q:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Vf1FrLWTFeY:R2K50ohef9Q:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Vf1FrLWTFeY:R2K50ohef9Q:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/Vf1FrLWTFeY" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/poor-cute-bunnies-likely-to-get-eaten-when-the-snow-melts-early/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/poor-cute-bunnies-likely-to-get-eaten-when-the-snow-melts-early/</feedburner:origLink></item>
		<item>
		<title>Stressed out moms mean faster growing babies—if you’re a squirrel</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/eRWOV8Z6Qw0/</link>
		<comments>http://arstechnica.com/science/2013/04/stressed-out-moms-mean-faster-growing-babiesif-youre-a-squirrel/#comments</comments>
		<pubDate>Thu, 18 Apr 2013 20:30:06 +0000</pubDate>
		<dc:creator>Kate Shaw</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[behavior]]></category>
		<category><![CDATA[Biology]]></category>
		<category><![CDATA[stress]]></category>
		<category><![CDATA[zoology]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=248589</guid>
		<description><![CDATA[For the offspring, it's grow fast, die young.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/red_squirrel-640x434.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://wdfw.wa.gov/conservation/gray_squirrel/">Washington State</a>
				</div>
	  </div>
  </div>
 <p>To give her offspring better odds in life, a mom can contribute all sorts of advantages: good genes, healthy milk, protection from predators, and more. Red squirrels—cousins of the more robust gray squirrels—are no exception. But new research in the journal <em>Science</em> shows that for these rodents, one such maternal boost stems from a surprising source: stress.</p>
<p>The research was conducted by a group of scientists who have spent the last 22 years studying a group of wild red squirrels in the Yukon. The ecosystem there follows an episodic pattern: every few years, when the spruce trees produce seeds en masse, <a href="http://arstechnica.com/science/2006/12/6385/">the squirrel population booms</a> in response. The team’s previous research had shown that in the years when squirrel numbers are especially high, fast-growing offspring tend to fare better than those that grow more slowly, presumably since they are better equipped to outcompete their rivals.</p>
<p>But how, exactly, do some squirrels end up growing so quickly?</p>
</div><p><a href="http://arstechnica.com/science/2013/04/stressed-out-moms-mean-faster-growing-babiesif-youre-a-squirrel/#p3">Read 9 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/stressed-out-moms-mean-faster-growing-babiesif-youre-a-squirrel/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eRWOV8Z6Qw0:qOA7FjLQ3-E:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=eRWOV8Z6Qw0:qOA7FjLQ3-E:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eRWOV8Z6Qw0:qOA7FjLQ3-E:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=eRWOV8Z6Qw0:qOA7FjLQ3-E:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eRWOV8Z6Qw0:qOA7FjLQ3-E:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eRWOV8Z6Qw0:qOA7FjLQ3-E:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/eRWOV8Z6Qw0" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/stressed-out-moms-mean-faster-growing-babiesif-youre-a-squirrel/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/stressed-out-moms-mean-faster-growing-babiesif-youre-a-squirrel/</feedburner:origLink></item>
		<item>
		<title>Kepler spots two super Earths, one squarely in the habitable zone</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/CmTcBzjd1Z8/</link>
		<comments>http://arstechnica.com/science/2013/04/kepler-spots-two-super-earths-one-squarely-in-the-habitable-zone/#comments</comments>
		<pubDate>Thu, 18 Apr 2013 18:30:15 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[astronomy]]></category>
		<category><![CDATA[exoplanets]]></category>
		<category><![CDATA[habitable zone]]></category>
		<category><![CDATA[kepler]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=248675</guid>
		<description><![CDATA[Planets are about 1.5 times Earth's radius—and the right distance from a host star.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/lores-640x455.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.cfa.harvard.edu/news/2013/pr201311.html">Harvard</a>
				</div>
	  </div>
  </div>
 <p>Liquid water is one of the essential ingredients of life on Earth and it has properties conducive to life in general, so researchers have defined the concept of a habitable zone with water as its focus. For astronomers, the habitable zone is the area far enough from a host star that a planet would be cool enough to support liquid water, but not so far that the water would be frozen. Although there are <a href="http://arstechnica.com/science/2013/02/the-ars-guide-to-exoplanets/">many caveats</a> to this definition—everything from the presence of greenhouse gasses to clouds will shift a planet's average surface temperature—the concept helps drive the search for a planet capable of supporting life.</p>
<p>The caveats, however, have led to a number of arguments over whether a given exoplanet is likely able to host liquid water, with some candidates shifting in and out of the habitable zone more than once. Still, it's pretty clear that given the <a href="http://arstechnica.com/science/2012/12/possible-habitable-zone-planet-a-mere-12-light-years-away/">large</a> <a href="http://arstechnica.com/science/2011/12/kepler-planet-candidate-count-clears-2000-first-habitable-zone-candidate-confirmed/">number</a> of <a href="http://arstechnica.com/science/2013/01/15-potential-planets-in-habitable-zones-found-by-citizen-scientists/">reported</a> <a href="http://arstechnica.com/science/2012/11/old-data-reveals-super-earth-lurking-in-a-nearby-stars-habitable-zone/">candidate</a> <a href="http://arstechnica.com/science/2007/04/aheavy-earth-like-planet-found-in-the-habitable-zone-of-a-red-dwarf/">exoplanets</a>, the odds suggest we've already spotted one. Today, scientists are announcing an exosolar system that has two planets that are both likely to be within the habitable zone, along with three others that are closer to the host star.</p>
<p>The findings come from the Kepler mission, which is dealing with a backlog of roughly 2,000 planet candidates to sort through and confirm. Normally, confirmation requires observing changes in the light emitted by the host star, which gets dragged closer to and further from the Earth as the planets circle it. However, this method requires that the planets either be large or orbiting close to the star so that their gravitational influence is big enough to shift it.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/kepler-spots-two-super-earths-one-squarely-in-the-habitable-zone/#p3">Read 7 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/kepler-spots-two-super-earths-one-squarely-in-the-habitable-zone/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=CmTcBzjd1Z8:7PyjiCACsZo:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=CmTcBzjd1Z8:7PyjiCACsZo:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=CmTcBzjd1Z8:7PyjiCACsZo:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=CmTcBzjd1Z8:7PyjiCACsZo:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=CmTcBzjd1Z8:7PyjiCACsZo:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=CmTcBzjd1Z8:7PyjiCACsZo:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/CmTcBzjd1Z8" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/kepler-spots-two-super-earths-one-squarely-in-the-habitable-zone/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/kepler-spots-two-super-earths-one-squarely-in-the-habitable-zone/</feedburner:origLink></item>
		<item>
		<title>Life is full of rude awakenings: all systems not go for Antares launch</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/fKSlkAe2kvI/</link>
		<comments>http://arstechnica.com/science/2013/04/life-is-full-of-rude-awakenings-all-systems-not-go-for-antares-launch/#comments</comments>
		<pubDate>Wed, 17 Apr 2013 22:25:18 +0000</pubDate>
		<dc:creator>Matthew Francis</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Antares rocket]]></category>
		<category><![CDATA[Commercial space]]></category>
		<category><![CDATA[NASA Wallops]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=248147</guid>
		<description><![CDATA[Premature detachment of an umbilical cable on Antares rocket means 48-hour delay.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/Antares_on_pad.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">The Antares rocket on its launchpad at NASA Wallops.</div>
	
			<div class="caption-credit">
							Matthew R. Francis				</div>
	  </div>
  </div>
 <p>As your reporter on the scene at NASA's Wallops Flight Facility, I intended to file an account of the test launch of the Antares rocket today. However, as things happen, the launch was aborted at the T-minus twelve minute mark, when a cable connected to the rocket's second stage disconnected prematurely.</p>
<p>Not only was this the first test of <a href="http://www.orbital.com/">Orbital Science's</a> Antares rocket, it was also to be the first launch from the new <a href="http://www.marsspaceport.com/">Mid-Atlantic Regional Spaceport (MARS)</a> launchpad at Wallops. As Orbital spokesman Barry Boneski said shortly after the decision to abort the launch, that means it's expected for things to go wrong. The disconnected cable, known as an umbilical, is a relatively minor problem, especially compared with some spectacularly failed tests of the past.</p>
<p>Boneski declined to speculate why the cable disconnected before its proper time. As he noted, the wind speed at the launchpad reached 20 knots (23 miles per hour, or 27 kilometers per hour) at one point, which may have played a role. However, Orbital and NASA crews had not yet had time to investigate fully.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/life-is-full-of-rude-awakenings-all-systems-not-go-for-antares-launch/#p3">Read 1 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/life-is-full-of-rude-awakenings-all-systems-not-go-for-antares-launch/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=fKSlkAe2kvI:6E-rGokYtHw:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=fKSlkAe2kvI:6E-rGokYtHw:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=fKSlkAe2kvI:6E-rGokYtHw:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=fKSlkAe2kvI:6E-rGokYtHw:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=fKSlkAe2kvI:6E-rGokYtHw:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=fKSlkAe2kvI:6E-rGokYtHw:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/fKSlkAe2kvI" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/life-is-full-of-rude-awakenings-all-systems-not-go-for-antares-launch/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/life-is-full-of-rude-awakenings-all-systems-not-go-for-antares-launch/</feedburner:origLink></item>
		<item>
		<title>Join Ars in watching the first attempted Antares rocket launch, 5pm EDT</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/TpY5YTV1CUg/</link>
		<comments>http://arstechnica.com/science/2013/04/join-ars-in-watching-the-first-attempted-antares-rocket-launch-5pm-edt/#comments</comments>
		<pubDate>Wed, 17 Apr 2013 20:00:51 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[antares]]></category>
		<category><![CDATA[NASA]]></category>
		<category><![CDATA[Orbital Science]]></category>
		<category><![CDATA[rocketry]]></category>
		<category><![CDATA[space]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247853</guid>
		<description><![CDATA[Ars is on the scene for future reports; for now, the livestream is here.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>As we <a href="http://arstechnica.com/science/2013/04/nasas-wallops-flight-facility-hosts-test-of-new-commercial-rocket/">mentioned yesterday</a>, today will mark Orbital Science's first attempt to send its Antares rocket into space. Meant to provide an alternative route to deliver supplies to the International Space Station, the Antares is ready for liftoff at NASA's Wallops facility on the east coast of Virginia.</p>
<p>Our own Matthew Francis is on hand to get pictures of the facility and liftoff attempt, but you can watch the events live on NASA TV. To save you the hassle of finding NASA TV, we've embedded the livestream below. Orbital Sciences has a lot of experience in the rocketry business, so we're crossing our fingers that its first liftoff with the new hardware will go off without a hitch.</p>
<p><iframe style="border: 0px none transparent;" src="http://www.ustream.tv/embed/6540154?v=3&amp;wmode=direct" height="302" width="480" frameborder="0" scrolling="no"></iframe></p>
</div><p><a href="http://arstechnica.com/science/2013/04/join-ars-in-watching-the-first-attempted-antares-rocket-launch-5pm-edt/#p3">Read 1 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/join-ars-in-watching-the-first-attempted-antares-rocket-launch-5pm-edt/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=TpY5YTV1CUg:0I7fIbU5hB4:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=TpY5YTV1CUg:0I7fIbU5hB4:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=TpY5YTV1CUg:0I7fIbU5hB4:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=TpY5YTV1CUg:0I7fIbU5hB4:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=TpY5YTV1CUg:0I7fIbU5hB4:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=TpY5YTV1CUg:0I7fIbU5hB4:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/TpY5YTV1CUg" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/join-ars-in-watching-the-first-attempted-antares-rocket-launch-5pm-edt/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/join-ars-in-watching-the-first-attempted-antares-rocket-launch-5pm-edt/</feedburner:origLink></item>
		<item>
		<title>Origin of life researchers figure out how to build bigger RNAs</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/Rp8HR3JnsAk/</link>
		<comments>http://arstechnica.com/science/2013/04/origin-of-life-researchers-figure-out-how-to-build-bigger-rnas/#comments</comments>
		<pubDate>Wed, 17 Apr 2013 18:00:15 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[biochemistry]]></category>
		<category><![CDATA[Biology]]></category>
		<category><![CDATA[chemistry]]></category>
		<category><![CDATA[origin of life]]></category>
		<category><![CDATA[RNA]]></category>
		<category><![CDATA[RNA world]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247831</guid>
		<description><![CDATA[Pre-biochemistry can link small building blocks into bigger molecules.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/groupI_intron_zoom-640x449.png"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">A close up of the active site of a catalytic RNA.</div>
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://web.chem.ucsb.edu/~molvisual/ribozymes.html">UCB</a>
				</div>
	  </div>
  </div>
 <p>We'll probably never know exactly how life on Earth got its start. The conditions in which it began have long since been lost, and there are simply too many precursor molecules and potential environments that could have gotten the process going. Nevertheless, researchers hope to put together a pathway that's at least plausible, starting from simple molecules that were present on the early Earth and building up to an enclosed system with basic inheritance (from there, evolution can take over).</p>
<p>A lot of progress has been made in understanding how a simple chemical, like <a href="http://arstechnica.com/science/2012/10/simple-reaction-makes-the-building-blocks-of-a-nucleic-acid/">hydrogen cyanide</a>, can be <a href="http://arstechnica.com/science/2009/05/origin-of-life-building-an-rna-world-from-simple-chemicals/">built up</a> through a series of reactions into a nucleotide, the basic building block of molecules like DNA and RNA. And we've learned quite a bit about how larger RNAs (more than 100 nucleotides long) can fold into complex structures that can catalyze reactions and undergo the chemical equivalent of Darwinian evolution. The challenge has been bridging the gap between the two, going from a handful of linked nucleotides to a large molecule that's potentially capable of catalyzing chemical reactions.</p>
<p>Now, the team that developed the earlier results is back with another publication. Their latest work shows how short molecules that are composed of just a handful of nucleotides can be linked together, eventually building longer, more complex chains. Once again, the chemistry is simple enough to occur on the early Earth, and the reaction might explain a curious bias in how DNA and RNA are built into long chains of nucleotides.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/origin-of-life-researchers-figure-out-how-to-build-bigger-rnas/#p3">Read 8 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/origin-of-life-researchers-figure-out-how-to-build-bigger-rnas/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Rp8HR3JnsAk:bFSI2-4KJes:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=Rp8HR3JnsAk:bFSI2-4KJes:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Rp8HR3JnsAk:bFSI2-4KJes:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=Rp8HR3JnsAk:bFSI2-4KJes:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Rp8HR3JnsAk:bFSI2-4KJes:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Rp8HR3JnsAk:bFSI2-4KJes:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/Rp8HR3JnsAk" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/origin-of-life-researchers-figure-out-how-to-build-bigger-rnas/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/origin-of-life-researchers-figure-out-how-to-build-bigger-rnas/</feedburner:origLink></item>
		<item>
		<title>Enzymes turn inedible waste plant material into food and fuel</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/7cIzKBjBCZg/</link>
		<comments>http://arstechnica.com/science/2013/04/enzymes-turn-inedible-waste-plant-material-into-food-and-fuel/#comments</comments>
		<pubDate>Wed, 17 Apr 2013 15:52:11 +0000</pubDate>
		<dc:creator>WIRED UK</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[biofuels]]></category>
		<category><![CDATA[cellulose]]></category>
		<category><![CDATA[ethanol]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247693</guid>
		<description><![CDATA[From cellulose to glucose to complex starches.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/wheatfield-640x430.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">Sunrise over a wheat field.</div>
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.flickr.com/photos/theknowlesgallery/4898802278/sizes/l/in/photostream/">The Knowles Gallery</a>
				</div>
	  </div>
  </div>
 <p>Researchers have managed to <a href="http://www.vtnews.vt.edu/articles/2013/04/041613-cals-zhangstarch.html">turn indigestible cellulose into starch</a>, a process that could render billions of tons of agricultural waste into food and fuel.</p>
<p>Plants grow more than 160 billion tons of cellulose—the material that makes up the walls of plant cells—every year, but only a tiny fraction of that is useful to humans in the crops we grow. This is frustrating, as cellulose is made up of glucose chains that are almost, but not quite, the same as those that make up the starch that constitutes 20 to 40 percent of most peoples' daily calorie intake.</p>
<p>With the world's population <a href="http://www.un.org/esa/population/publications/longrange2/WorldPop2300final.pdf">forecast</a> to reach nine billion by 2050, working out how to alter cellulose glucose into something more practical could be vital for preventing starvation. There's also an extra benefit in that some could be used for biofuels.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/enzymes-turn-inedible-waste-plant-material-into-food-and-fuel/#p3">Read 8 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/enzymes-turn-inedible-waste-plant-material-into-food-and-fuel/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=7cIzKBjBCZg:wXuUxHirU7o:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=7cIzKBjBCZg:wXuUxHirU7o:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=7cIzKBjBCZg:wXuUxHirU7o:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=7cIzKBjBCZg:wXuUxHirU7o:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=7cIzKBjBCZg:wXuUxHirU7o:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=7cIzKBjBCZg:wXuUxHirU7o:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/7cIzKBjBCZg" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/enzymes-turn-inedible-waste-plant-material-into-food-and-fuel/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/enzymes-turn-inedible-waste-plant-material-into-food-and-fuel/</feedburner:origLink></item>
		<item>
		<title>Microsoft Excel: The ruiner of global economies?</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/f4V4lDXsTo4/</link>
		<comments>http://arstechnica.com/tech-policy/2013/04/microsoft-excel-the-ruiner-of-global-economies/#comments</comments>
		<pubDate>Tue, 16 Apr 2013 22:20:08 +0000</pubDate>
		<dc:creator>Peter Bright</dc:creator>
				<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Scientific Method]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247197</guid>
		<description><![CDATA[A paper used to justify austerity economics appears to contain an Excel error.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/excel-destroys-earth-640x360.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">Figuratively speaking, of course.</div>
	
			<div class="caption-credit">
							Aurich Lawson				</div>
	  </div>
  </div>
 <p>An economics paper claiming that high levels of national debt led to low or negative economic growth could turn out to be deeply flawed as a result of, among other things, an incorrect formula in an Excel spreadsheet. Microsoft's PowerPoint has been <a href="http://www.wired.com/wired/archive/11.09/ppt2.html">considered evil</a> thanks to the proliferation of poorly presented data and dull slides that are created with it. Might Excel also deserve such hyperbolic censure?</p>
<p>The paper, <a href="http://www.nber.org/papers/w15639.pdf">Growth in a Time of Debt</a>, was written by economists Carmen Reinhart and Kenneth Rogoff and published in 2010. Since publication, it has been cited <a href="http://www.reinhartandrogoff.com/related-research/growth-in-a-time-of-debt-featured-in">abundantly</a> by the world's press politicians, including one-time vice president nominee Paul Ryan (R-WI). The link it draws between high levels of debt and negative average economic growth has been used by right-leaning politicians to justify austerity budgets: slashing government expenditure and reducing budget deficits in a bid to curtail the growth of debt.</p>
<p>This link was always controversial, with many economists proposing that the correlation between high debt and low growth was just as likely to have a causal link in the other direction to that proposed by Reinhart and Rogoff: it's not that high debt causes low growth, but rather that low growth leads to high debt.</p>
</div><p><a href="http://arstechnica.com/tech-policy/2013/04/microsoft-excel-the-ruiner-of-global-economies/#p3">Read 10 remaining paragraphs</a> | <a href="http://arstechnica.com/tech-policy/2013/04/microsoft-excel-the-ruiner-of-global-economies/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=f4V4lDXsTo4:5KqO8QIKL5s:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=f4V4lDXsTo4:5KqO8QIKL5s:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=f4V4lDXsTo4:5KqO8QIKL5s:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=f4V4lDXsTo4:5KqO8QIKL5s:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=f4V4lDXsTo4:5KqO8QIKL5s:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=f4V4lDXsTo4:5KqO8QIKL5s:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/f4V4lDXsTo4" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/tech-policy/2013/04/microsoft-excel-the-ruiner-of-global-economies/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/tech-policy/2013/04/microsoft-excel-the-ruiner-of-global-economies/</feedburner:origLink></item>
		<item>
		<title>NASA’s Wallops Flight Facility hosts test of new commercial rocket</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/wCNxbCjC5Q8/</link>
		<comments>http://arstechnica.com/science/2013/04/nasas-wallops-flight-facility-hosts-test-of-new-commercial-rocket/#comments</comments>
		<pubDate>Tue, 16 Apr 2013 20:10:10 +0000</pubDate>
		<dc:creator>Matthew Francis</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Antares rocket]]></category>
		<category><![CDATA[Commercial space]]></category>
		<category><![CDATA[NASA]]></category>
		<category><![CDATA[NASA Wallops]]></category>
		<category><![CDATA[orbital sciences]]></category>
		<category><![CDATA[rocket science]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=247013</guid>
		<description><![CDATA[Antares rocket from Orbital Sciences could serve International Space Station.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/Antares2.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">The new Antares rocket will be test-launched from NASA Wallops this week.</div>
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.nasa.gov/exploration/commercial/cargo/orbitalsciences-index.html">NASA/Brea Reeves</a>
				</div>
	  </div>
  </div>
 <p>For decades, the Space Shuttles were the primary way the United States transported cargo and people into low-Earth orbit. Upon retirement of the aging fleet of Shuttles, NASA has promoted a public-private partnership with companies that are developing new rockets for transport into low-Earth orbit. One of these companies, <a href="http://www.orbital.com/">Orbital Sciences Corporation</a>, is test-launching a new rocket design at NASA's <a href="http://www.nasa.gov/centers/wallops/home/index.html">Wallops Flight Facility</a> in Virginia this week.</p>
<p>Orbital's design, known as Antares, is a multi-purpose middleweight rocket built to carry non-human payloads into space. A major part of its intended purview is robotic delivery to the International Space Station (ISS), an important task for supplying long-duration stays. This week's launch from Wallops is the first orbital test of Antares.</p>
<p>While NASA's Commercial Orbital Transportation Services (COTS) program for Station resupply only dates to 2006, Orbital has been developing rockets since 1990, beginning with the airplane-launched Pegasus. (For comparison, SpaceX was founded in 2002 and Virgin Galactic began operations in 2004.) Antares is based on Orbital's earlier designs, including the ground-launched Minotaur and Taurus rockets. However, the new design is much heftier, capable of carrying up to 6,120 kilograms (about 13,500 pounds), much higher than Minotaur's 1,730 kg maximum. According to the Orbital website, Antares is also capable of launching payloads on interplanetary trajectories.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/nasas-wallops-flight-facility-hosts-test-of-new-commercial-rocket/#p3">Read 4 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/nasas-wallops-flight-facility-hosts-test-of-new-commercial-rocket/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=wCNxbCjC5Q8:sMcSNGlZhXw:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=wCNxbCjC5Q8:sMcSNGlZhXw:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=wCNxbCjC5Q8:sMcSNGlZhXw:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=wCNxbCjC5Q8:sMcSNGlZhXw:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=wCNxbCjC5Q8:sMcSNGlZhXw:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=wCNxbCjC5Q8:sMcSNGlZhXw:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/wCNxbCjC5Q8" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/nasas-wallops-flight-facility-hosts-test-of-new-commercial-rocket/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/nasas-wallops-flight-facility-hosts-test-of-new-commercial-rocket/</feedburner:origLink></item>
		<item>
		<title>Latest dark matter search comes a bit short of discovery</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/yPVyWMogmo8/</link>
		<comments>http://arstechnica.com/science/2013/04/latest-dark-matter-search-comes-a-bit-short-of-discovery/#comments</comments>
		<pubDate>Tue, 16 Apr 2013 18:30:55 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Dark Matter]]></category>
		<category><![CDATA[particle physics]]></category>
		<category><![CDATA[Physics]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=246871</guid>
		<description><![CDATA[But the particles it sees are lighter than expected.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/tamu_detector_updated.jpg"><div class="caption" style="font-size:0.8em">
			<div class="caption-text">One of the detectors that will be used in the next iteration of this hunt for dark matter.</div>
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.science.tamu.edu/articles/1052">Texas A&amp;M</a>
				</div>
	  </div>
  </div>
 <p>The evidence for dark matter is comprehensive. We've measured its effect on galaxies and galaxy clusters, and we've seen its imprint in gravitational lenses and the cosmic microwave background. The annoying thing is that we still don't know <em>what it is.</em> All the evidence indicates that dark matter is likely to be a weakly interacting massive particle (or WIMP), but the best we've found when searching for this particle is a few intriguing hints of evidence.</p>
<p>In the latest hint, the people behind a detector that sits deep in a Minnesota mine say they've observed three events that appear likely to be the faint impact of dark matter particles. Unfortunately, by every measure they try, the significance of these events falls short of discovery. We're therefore left waiting for better detectors and more data—just as we were the last time this team <a href="http://arstechnica.com/science/2009/12/two-events-hint-of-the-impact-of-dark-matter-particles/">announced a similar result</a>.</p>
<p>Since the evidence indicates that dark matter is a particle (after all, it had to be around before there were even atoms in order to influence the cosmic microwave background), researchers have focused on three ways we might detect a particle that doesn't fit into the existing Standard Model. One track involves searches through the data in the LHC, looking for signs that some of the mass produced in a collision is being carried away by something we can't detect. So far, that hunt has come up blank.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/latest-dark-matter-search-comes-a-bit-short-of-discovery/#p3">Read 8 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/latest-dark-matter-search-comes-a-bit-short-of-discovery/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=yPVyWMogmo8:a20vWDiwoAY:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=yPVyWMogmo8:a20vWDiwoAY:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=yPVyWMogmo8:a20vWDiwoAY:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=yPVyWMogmo8:a20vWDiwoAY:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=yPVyWMogmo8:a20vWDiwoAY:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=yPVyWMogmo8:a20vWDiwoAY:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/yPVyWMogmo8" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/latest-dark-matter-search-comes-a-bit-short-of-discovery/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/latest-dark-matter-search-comes-a-bit-short-of-discovery/</feedburner:origLink></item>
		<item>
		<title>Is the Grand Canyon showing its age?</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/O4WOsv5fnJc/</link>
		<comments>http://arstechnica.com/science/2013/04/is-the-grand-canyon-showing-its-age/#comments</comments>
		<pubDate>Tue, 16 Apr 2013 15:00:35 +0000</pubDate>
		<dc:creator>Ars Staff</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Earth science]]></category>
		<category><![CDATA[geology]]></category>
		<category><![CDATA[grand canyon]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=246597</guid>
		<description><![CDATA[Researchers argue over whether the Canyon is literally as old as the dinosaurs.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/20070921161701-640x422.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.nps.gov/grca/photosmultimedia/GRAND-CANYON-SCENIC-SPLENDOR.htm">National Park Service</a>
				</div>
	  </div>
  </div>
 <p>It might be one of the best studied geologic features in the country, and it's certainly one of the most striking, but a scientific debate continues to rage over the Grand Canyon. Specifically over what's typically considered an impolite question: how old is it?</p>
<p>Many geologists estimate that the Colorado River began to carve the canyon we see today between five and six million years ago. But in December, a paper published in <i>Science</i> presented evidence that the western portion of the canyon might have been carved close to the present depth by 70 million years ago.</p>
<p>This week, <i>Science</i> hosts two comments challenging the ancient canyon proposal and a defense of the controversial paper from its authors. This is science in action—researchers debating a 65 million year conflict between evidence from traditional geological methods and new dating technology.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/is-the-grand-canyon-showing-its-age/#p3">Read 8 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/is-the-grand-canyon-showing-its-age/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=O4WOsv5fnJc:VSpFFzS5Fjk:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=O4WOsv5fnJc:VSpFFzS5Fjk:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=O4WOsv5fnJc:VSpFFzS5Fjk:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=O4WOsv5fnJc:VSpFFzS5Fjk:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=O4WOsv5fnJc:VSpFFzS5Fjk:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=O4WOsv5fnJc:VSpFFzS5Fjk:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/O4WOsv5fnJc" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/is-the-grand-canyon-showing-its-age/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/is-the-grand-canyon-showing-its-age/</feedburner:origLink></item>
		<item>
		<title>Supreme Court debates whether “snipping” human genes merits a patent</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/VcfzCOtCvXw/</link>
		<comments>http://arstechnica.com/tech-policy/2013/04/supreme-court-debates-whether-snipping-human-genes-merits-a-patent/#comments</comments>
		<pubDate>Mon, 15 Apr 2013 22:15:46 +0000</pubDate>
		<dc:creator>Timothy B. Lee</dc:creator>
				<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[American Civil Liberties Union]]></category>
		<category><![CDATA[Myriad Genetics]]></category>
		<category><![CDATA[patent law]]></category>
		<category><![CDATA[Supreme Court]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=246247</guid>
		<description><![CDATA[Some justices are skeptical of the "very patent happy" patent office.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/3210447964_eaf2399202_b-640x425.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.flickr.com/photos/mikep1979/3210447964/sizes/l/in/photostream/">Mike Paradise</a>
				</div>
	  </div>
  </div>
 <p>The word of the day at the Supreme Court on Monday was "snip." The biotech company Myriad Genetics is defending <a href="http://arstechnica.com/tech-policy/2013/04/will-the-supreme-court-end-human-gene-patents-after-three-decades/">patents that give it exclusive control over two genes</a> linked to breast cancer. Critics have argued that Myriad merely "snipped" the genes from the human genome, and that this no more deserves patent protection than "snipping" a leaf from a plant or removing a liver from a human body.</p>
<p>"Here, what's involved is snipping," Chief Justice John Roberts <a href="http://www.supremecourt.gov/oral_arguments/argument_transcripts/12-398-amc7.pdf">said</a> to Myriad's lawyer, Gregory Castanias. "You've got the thing there and you snip off the top and you snip off the bottom and there you've got it."</p>
<p>"You wouldn't even know where to snip until the Myriad invention," Castanias countered.</p>
</div><p><a href="http://arstechnica.com/tech-policy/2013/04/supreme-court-debates-whether-snipping-human-genes-merits-a-patent/#p3">Read 14 remaining paragraphs</a> | <a href="http://arstechnica.com/tech-policy/2013/04/supreme-court-debates-whether-snipping-human-genes-merits-a-patent/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=VcfzCOtCvXw:9cKxFvmGh1c:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=VcfzCOtCvXw:9cKxFvmGh1c:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=VcfzCOtCvXw:9cKxFvmGh1c:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=VcfzCOtCvXw:9cKxFvmGh1c:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=VcfzCOtCvXw:9cKxFvmGh1c:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=VcfzCOtCvXw:9cKxFvmGh1c:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/VcfzCOtCvXw" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/tech-policy/2013/04/supreme-court-debates-whether-snipping-human-genes-merits-a-patent/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/tech-policy/2013/04/supreme-court-debates-whether-snipping-human-genes-merits-a-patent/</feedburner:origLink></item>
		<item>
		<title>Making HIV tests visible to the naked eye</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/LQZUTAQcvz4/</link>
		<comments>http://arstechnica.com/science/2013/04/making-hiv-tests-visible-to-the-naked-eye/#comments</comments>
		<pubDate>Mon, 15 Apr 2013 12:54:12 +0000</pubDate>
		<dc:creator>Chris Lee</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[biochemistry]]></category>
		<category><![CDATA[Biology]]></category>
		<category><![CDATA[chemistry]]></category>
		<category><![CDATA[developing world]]></category>
		<category><![CDATA[healthcare]]></category>
		<category><![CDATA[HIV]]></category>
		<category><![CDATA[medicine]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=244705</guid>
		<description><![CDATA[Pregnancy-test-style diagnosis gives the blue stripe a whole new meaning.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/pregnancy_test_positive.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.personal.psu.edu/afr3/blogs/siowfa12/2012/12/pregnancy-tests.html">Penn State</a>
				</div>
	  </div>
  </div>
 <p>One of the recurring themes in healthcare is that a lot of what works in the developed world doesn't work in the developing world. There may be limited access to basics like power or communications facilities, never mind the medical technologies that make a hospital one of the modern wonders of the world. Beyond simple access to expensive technologies and medicines, even something as basic as a diagnostic test might be too expensive or require skilled technicians to use.</p>
<p>Though cures might remain expensive—drug companies like their profits—every dime saved on diagnosis is a dime more for prevention and cure. That makes cheap, accurate, and simple diagnostic tests very, very desirable.</p>
<p>I was thinking about this when I came across an older paper that somehow didn't attract any attention when it came out. Last year, a group of researchers <a href="http://dx.doi.org/10.1038/nnano.2012.18">showed</a> that they could detect HIV at extremely low concentrations. That by itself is nothing special: people are always improving diagnostic tests. What is special is that the test is very much like a pregnancy test, in that a simple visible color change indicates a positive result. Even better, it seems to work in real-life tests.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/making-hiv-tests-visible-to-the-naked-eye/#p3">Read 11 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/making-hiv-tests-visible-to-the-naked-eye/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=LQZUTAQcvz4:Yu1dzk3Ur78:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=LQZUTAQcvz4:Yu1dzk3Ur78:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=LQZUTAQcvz4:Yu1dzk3Ur78:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=LQZUTAQcvz4:Yu1dzk3Ur78:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=LQZUTAQcvz4:Yu1dzk3Ur78:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=LQZUTAQcvz4:Yu1dzk3Ur78:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/LQZUTAQcvz4" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/making-hiv-tests-visible-to-the-naked-eye/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/making-hiv-tests-visible-to-the-naked-eye/</feedburner:origLink></item>
		<item>
		<title>New F-1B rocket engine upgrades Apollo-era design with 1.8M lbs of thrust</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/ftrk7zc61xk/</link>
		<comments>http://arstechnica.com/science/2013/04/new-f-1b-rocket-engine-upgrades-apollo-era-deisgn-with-1-8m-lbs-of-thrust/#comments</comments>
		<pubDate>Mon, 15 Apr 2013 01:00:30 +0000</pubDate>
		<dc:creator>Lee Hutchinson</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[dynetics]]></category>
		<category><![CDATA[f-1]]></category>
		<category><![CDATA[F-1B]]></category>
		<category><![CDATA[NASA]]></category>
		<category><![CDATA[pratt whitney rocketdyne]]></category>
		<category><![CDATA[pwr]]></category>
		<category><![CDATA[rocketdybe]]></category>
		<category><![CDATA[rockets]]></category>
		<category><![CDATA[sls]]></category>
		<category><![CDATA[space]]></category>
		<category><![CDATA[space flight]]></category>
		<category><![CDATA[space launch system]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=243439</guid>
		<description><![CDATA[Dynetics and Pratt Whitney Rocketdyne rebuild the F-1 for the "Pyrios" booster.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>NASA has spent a lot of time and money <a href="http://arstechnica.com/science/2013/01/saturn-v-moon-rocket-engine-firing-again-after-40-years-sort-of/">resurrecting the F-1 rocket engine</a> that powered the Saturn V back in the 1960s and 1970s, and Ars recently spent a week at the Marshall Space Flight Center in Huntsville, Alabama, to get <a href="http://arstechnica.com/science/2013/04/how-nasa-brought-the-monstrous-f-1-moon-rocket-back-to-life/">the inside scoop on how the effort came to be</a>. But there's a very practical reason why NASA is putting old rocket parts up on a test stand and firing them off: its latest launch vehicle might be powered by engines that look, sound, and work a whole lot like the legendary F-1.</p>
<p>This new launch vehicle, known as the <a href="https://en.wikipedia.org/wiki/Space_Launch_System">Space Launch System</a>, or SLS, is currently taking shape on NASA drawing boards. However, as is its mandate, NASA won't be <em>building</em> the rocket itself—it will allow private industry to bid for the rights to build various components. One potential design wrinkle in SLS is that instead of using Space Shuttle-style solid rocket boosters, SLS could instead use liquid-fueled rocket motors, which would make it the United States' first human-rated rocket in more than 30 years not to use solid-fuel boosters.</p>
<p>The contest to suss this out is the <a href="http://www.nasaspaceflight.com/2012/11/dynetics-pwr-liquidize-sls-booster-competition-f-1-power/">Advanced Booster Competition</a>, and one of the companies that has been down-selected as a final competitor is Huntsville-based <a href="http://www.dynetics.com">Dynetics</a>. Dynetics has partnered with Pratt Whitney Rocketdyne (designers of the Saturn V's F-1 engine, among others) to propose a liquid-fueled booster featuring an engine based heavily on the design of the famous F-1. The booster is tentatively named <em>Pyrios</em>, after one of the fiery horses that pulled the god Apollo's chariot; the engine is being called the F-1B.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/new-f-1b-rocket-engine-upgrades-apollo-era-deisgn-with-1-8m-lbs-of-thrust/#p3">Read 34 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/new-f-1b-rocket-engine-upgrades-apollo-era-deisgn-with-1-8m-lbs-of-thrust/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=ftrk7zc61xk:5ORR3r677Os:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=ftrk7zc61xk:5ORR3r677Os:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=ftrk7zc61xk:5ORR3r677Os:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=ftrk7zc61xk:5ORR3r677Os:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=ftrk7zc61xk:5ORR3r677Os:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=ftrk7zc61xk:5ORR3r677Os:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/ftrk7zc61xk" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/new-f-1b-rocket-engine-upgrades-apollo-era-deisgn-with-1-8m-lbs-of-thrust/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/new-f-1b-rocket-engine-upgrades-apollo-era-deisgn-with-1-8m-lbs-of-thrust/</feedburner:origLink></item>
		<item>
		<title>Gallery: Behind the scenes at NASA’s Marshall Space Flight Center</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/IdBEioMsgeE/</link>
		<comments>http://arstechnica.com/science/2013/04/gallery-behind-the-scenes-at-nasas-marshall-space-flight-center/#comments</comments>
		<pubDate>Mon, 15 Apr 2013 01:00:29 +0000</pubDate>
		<dc:creator>Lee Hutchinson</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[f-1]]></category>
		<category><![CDATA[F-1B]]></category>
		<category><![CDATA[gallery]]></category>
		<category><![CDATA[huntsville]]></category>
		<category><![CDATA[marshall space flight center]]></category>
		<category><![CDATA[msfc]]></category>
		<category><![CDATA[NASA]]></category>
		<category><![CDATA[rocketry]]></category>
		<category><![CDATA[Saturn V]]></category>
		<category><![CDATA[space]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=244771</guid>
		<description><![CDATA[We watched a rocket test and came away with a ton of awesome photographs.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>Ars spent the better part of a week at NASA's Marshall Space Flight Center (MSFC) in Huntsville, Alabama, gathering material for our pieces on the Saturn V's F-1 <a href="http://arstechnica.com/science/2013/04/how-nasa-brought-the-monstrous-f-1-moon-rocket-back-to-life/">resurrection and testing</a> and the <a href="http://arstechnica.com/science/2013/04/new-f-1b-rocket-engine-upgrades-apollo-era-deisgn-with-1-8m-lbs-of-thrust/">new F-1B</a> rocket engine that's being developed right now. In the process, we got to peek into a number of different areas of MSFC not open to the general public, including manufacturing, training, and operations areas. I took over a thousand photographs over the time I was there, and couldn't resist showing them off. Below are the 40 best images from our tour of Marshall, the birthplace of the Saturn V rocket and the place where NASA's <a href="https://en.wikipedia.org/wiki/Space_Launch_System">latest launch vehicle</a> is being designed.</p>
  <div class="gallery">
    <div class="gallery-main-image" style="width:auto">
      <div class="gallery-image-container" style="height:auto">
        <div class="gallery-image-wrap">
        <a class="enlarge" href="http://cdn.arstechnica.net/wp-content/uploads/2013/04/eande-ss-f1-moonrise.jpg" data-width="2048" data-height="1365">
          <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/eande-ss-f1-moonrise-980x653.jpg" style="max-width:auto;max-height:auto"></a>
        </div>
      </div>
    </div>

        <div class="gallery-image-description">

            <h2>F-1 moonrise</h2>
                <p>A late-afternoon moon peeking over the edge of the engine interface panel on one of Marshall Space Flight Center's outdoor F-1 rockets. </p>
      
          <p class="gallery-image-credit">
              <em>Lee Hutchinson</em>
            </p>
    
    </div>
    
       
          <p><strong>40 more images in gallery</strong></p>
      </div>
    <style type="text/css"> .related-stories { display: none !important; } </style>
</div><p><a href="http://arstechnica.com/science/2013/04/gallery-behind-the-scenes-at-nasas-marshall-space-flight-center/">Read on Ars Technica</a> | <a href="http://arstechnica.com/science/2013/04/gallery-behind-the-scenes-at-nasas-marshall-space-flight-center/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=IdBEioMsgeE:eo0nYce5DmY:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=IdBEioMsgeE:eo0nYce5DmY:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=IdBEioMsgeE:eo0nYce5DmY:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=IdBEioMsgeE:eo0nYce5DmY:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=IdBEioMsgeE:eo0nYce5DmY:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=IdBEioMsgeE:eo0nYce5DmY:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/IdBEioMsgeE" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/gallery-behind-the-scenes-at-nasas-marshall-space-flight-center/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/gallery-behind-the-scenes-at-nasas-marshall-space-flight-center/</feedburner:origLink></item>
		<item>
		<title>How NASA brought the monstrous F-1 “moon rocket” engine back to life</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/Ij4-OtGe614/</link>
		<comments>http://arstechnica.com/science/2013/04/how-nasa-brought-the-monstrous-f-1-moon-rocket-back-to-life/#comments</comments>
		<pubDate>Mon, 15 Apr 2013 01:00:03 +0000</pubDate>
		<dc:creator>Lee Hutchinson</dc:creator>
				<category><![CDATA[Features]]></category>
		<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[booster]]></category>
		<category><![CDATA[f-1]]></category>
		<category><![CDATA[F-1B]]></category>
		<category><![CDATA[kerosene]]></category>
		<category><![CDATA[marshall space flight center]]></category>
		<category><![CDATA[msfc]]></category>
		<category><![CDATA[NASA]]></category>
		<category><![CDATA[pratt and whitney]]></category>
		<category><![CDATA[rocket]]></category>
		<category><![CDATA[rocket engine]]></category>
		<category><![CDATA[rocketdyne]]></category>
		<category><![CDATA[Saturn V]]></category>
		<category><![CDATA[space]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=221211</guid>
		<description><![CDATA[The story of young engineers who resurrected an engine nearly twice their age.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>There has never been anything like the <a href="https://en.wikipedia.org/wiki/Saturn_V">Saturn V</a>, the launch vehicle that powered the United States past the Soviet Union to a series of manned lunar landings in the late 1960s and early 1970s. The rocket redefined "massive," standing 363 feet (110 meters) in height and producing a ludicrous 7.68 million pounds (34 meganewtons) of thrust from the five monstrous, kerosene-gulping <a href="https://en.wikipedia.org/wiki/F-1_%28rocket_engine%29">Rocketdyne F-1</a> rocket engines that made up its first stage.</p>
<p>At the time, the F-1 was the largest and most powerful liquid-fueled engine ever constructed; even today, its design remains unmatched (though see the sidebar, "The Soviets," for more information on engines that have rivaled the F-1). The power generated by five of these engines was best conceptualized by author David Woods in his book <em><a href="http://books.google.com/books?id=x-taL4N0sjIC&amp;lpg=PP1&amp;pg=PA79#v=onepage&amp;q&amp;f=false">How Apollo Flew to the Moon</a>—</em>"[T]he power output of the Saturn first stage was 60 gigawatts. This happens to be very similar to the peak electricity demand of the United Kingdom."</p>
<p>Despite the stunning success of the Saturn V, NASA's direction shifted after Project Apollo's conclusion; the <a href="http://arstechnica.com/science/2011/07/30-years-of-manned-space-flight-the-last-flight-of-the-space-shuttle/">Space Transport System</a>—the Space Shuttle and its associated hardware—was instead designed with wildly different engines. For thirty years, NASA's astronaut corps rode into orbit aboard Space Shuttles powered by <a href="https://en.wikipedia.org/wiki/RS-25">RS-25</a> liquid hydrogen-powered engines and <a href="https://en.wikipedia.org/wiki/Space_Shuttle_Solid_Rocket_Booster">solid-propellant boosters</a>. With the Shuttle's discontinuation, NASA is currently hitching space rides with the Russians.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/how-nasa-brought-the-monstrous-f-1-moon-rocket-back-to-life/#p3">Read 64 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/how-nasa-brought-the-monstrous-f-1-moon-rocket-back-to-life/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Ij4-OtGe614:LDXOJCO_W1c:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=Ij4-OtGe614:LDXOJCO_W1c:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Ij4-OtGe614:LDXOJCO_W1c:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=Ij4-OtGe614:LDXOJCO_W1c:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Ij4-OtGe614:LDXOJCO_W1c:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=Ij4-OtGe614:LDXOJCO_W1c:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/Ij4-OtGe614" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/how-nasa-brought-the-monstrous-f-1-moon-rocket-back-to-life/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/how-nasa-brought-the-monstrous-f-1-moon-rocket-back-to-life/</feedburner:origLink></item>
		<item>
		<title>Will the Supreme Court end human gene patents after three decades?</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/GxiUegKxh7M/</link>
		<comments>http://arstechnica.com/tech-policy/2013/04/will-the-supreme-court-end-human-gene-patents-after-three-decades/#comments</comments>
		<pubDate>Sun, 14 Apr 2013 20:00:29 +0000</pubDate>
		<dc:creator>Timothy B. Lee</dc:creator>
				<category><![CDATA[Law & Disorder]]></category>
		<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Gene patents]]></category>
		<category><![CDATA[Myriad Genetics]]></category>
		<category><![CDATA[patent]]></category>
		<category><![CDATA[Supreme Court]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=244687</guid>
		<description><![CDATA[Court considers invalidating breast cancer gene patents&#8212;and thousands of others.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/51784073_941f42961b_b-640x512.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.flickr.com/photos/memestate/51784073/">Rich Anderson</a>
				</div>
	  </div>
  </div>
 <p>Since the 1980s, patent lawyers have been claiming pieces of humanity's genetic code. The United States Patent and Trademark Office has granted thousands of gene patents. The Federal Circuit, the court that hears all patent appeals, has consistently ruled such patents are legal.</p>
<p>But the judicial winds have been shifting. The Supreme Court has never ruled on the legality of gene patents. And recently, the Supreme Court has grown increasingly skeptical of the Federal Circuit's <a href="http://arstechnica.com/tech-policy/2012/09/how-a-rogue-appeals-court-wrecked-the-patent-system/">patent-friendly jurisprudence.</a></p>
<p>Meanwhile, a growing number of researchers, health care providers, and public interest groups have raised concerns about the harms of gene patents. The American Civil Liberties Union <a href="http://www.aclu.org/blog/free-speech-technology-and-liberty-womens-rights/voices-human-gene-patents-its-time-free-our">estimates</a> that more than 40 percent of genes are now patented. Those patents have created "patent thickets" that make it difficult for scientists to do genetic research and commercialize their results. Monopolies on genetic testing have raised prices and reduced patient options.</p>
</div><p><a href="http://arstechnica.com/tech-policy/2013/04/will-the-supreme-court-end-human-gene-patents-after-three-decades/#p3">Read 14 remaining paragraphs</a> | <a href="http://arstechnica.com/tech-policy/2013/04/will-the-supreme-court-end-human-gene-patents-after-three-decades/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=GxiUegKxh7M:GLxQ8u7wM40:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=GxiUegKxh7M:GLxQ8u7wM40:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=GxiUegKxh7M:GLxQ8u7wM40:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=GxiUegKxh7M:GLxQ8u7wM40:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=GxiUegKxh7M:GLxQ8u7wM40:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=GxiUegKxh7M:GLxQ8u7wM40:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/GxiUegKxh7M" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/tech-policy/2013/04/will-the-supreme-court-end-human-gene-patents-after-three-decades/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/tech-policy/2013/04/will-the-supreme-court-end-human-gene-patents-after-three-decades/</feedburner:origLink></item>
		<item>
		<title>This AI “solves” Super Mario Bros. and other classic NES games</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/eP42pLcuB8U/</link>
		<comments>http://arstechnica.com/gaming/2013/04/this-ai-solves-super-mario-bros-and-other-classic-nes-games/#comments</comments>
		<pubDate>Sun, 14 Apr 2013 19:00:13 +0000</pubDate>
		<dc:creator>WIRED UK</dc:creator>
				<category><![CDATA[Opposable Thumbs]]></category>
		<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[AI]]></category>
		<category><![CDATA[Super Mario Bros.]]></category>
		<category><![CDATA[super nes]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=245545</guid>
		<description><![CDATA[Program even takes advantage of bugs and glitches.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/snes-640x476.png"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							Nintendo				</div>
	  </div>
  </div>
 <p><figure class="video" style="width:640px"><iframe style="display:block" type="text/html" width="640" height="360" src="http://www.youtube-nocookie.com/embed/xOCurBYI_gY?start=0&amp;wmode=transparent" frameborder="0"></iframe></figure><br>
In the 28 years since <i>Super Mario Bros.</i> was released, and it's obviously been comprehensively beaten, thoroughly, many thousands of times in that time by players around the world. But have you ever made the game beat itself?</p>
<p>That's what computer scientist Tom Murphy has done. At <a href="http://sigbovik.org/">SigBovik 2013</a>, he presented a program that "solves" how to play <i>Super Mario Bros</i>., or any other NES game, like it's just another kind of mathematical problem. And for those who know that SigBovik is an annual computer science conference dedicated to spoof research, hosted on April 1 every year, Murphy stresses that this is "100 percent real."</p>
<p>He outlines his method in a paper, "<a href="http://www.cs.cmu.edu/~tom7/mario/mario.pdf">The First Level of <i>Super Mario Bros.</i> is Easy with Lexicographic Orderings and Time Travel... after that it gets a little tricky</a>," but he also presented the results in the video you can see with this story.</p>
</div><p><a href="http://arstechnica.com/gaming/2013/04/this-ai-solves-super-mario-bros-and-other-classic-nes-games/#p3">Read 11 remaining paragraphs</a> | <a href="http://arstechnica.com/gaming/2013/04/this-ai-solves-super-mario-bros-and-other-classic-nes-games/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eP42pLcuB8U:JPNHxrFqAis:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=eP42pLcuB8U:JPNHxrFqAis:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eP42pLcuB8U:JPNHxrFqAis:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=eP42pLcuB8U:JPNHxrFqAis:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eP42pLcuB8U:JPNHxrFqAis:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=eP42pLcuB8U:JPNHxrFqAis:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/eP42pLcuB8U" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/gaming/2013/04/this-ai-solves-super-mario-bros-and-other-classic-nes-games/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/gaming/2013/04/this-ai-solves-super-mario-bros-and-other-classic-nes-games/</feedburner:origLink></item>
		<item>
		<title>Weird Science retracts its penis to keep everyone guessing</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/8jSlZ98435s/</link>
		<comments>http://arstechnica.com/science/2013/04/weird-science-retracts-its-penis-to-keep-everyone-guessing/#comments</comments>
		<pubDate>Sat, 13 Apr 2013 20:30:21 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Week in review]]></category>
		<category><![CDATA[Weird Science]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=245907</guid>
		<description><![CDATA[And, scientific achievement goes hand-in-hand with beer consumption. <em>Duh.</em>]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/4129761544_b4cec35a55_z.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							<a rel="nofollow" href="http://www.flickr.com/photos/photographicpictures/4129761544/">Chris Cotterman</a>
				</div>
	  </div>
  </div>
 <p><strong>Ever wonder about the consequences of your penis not being retractable?</strong> Wonder no more. Fortunately, a team of Australian researchers are <a href="http://www.pnas.org/content/early/2013/04/03/1219361110">willing to lay it out for us</a>: "Before clothing, the nonretractable human penis would have been conspicuous to potential mates." Since clothing is a relatively recent development, there's a chance that its conspicuousness made a mark on human evolution, either in terms of male anatomy or in terms of female tastes. The authors grabbed a database from a study of the Italian male population and used it to make computer-generated male bodies. They showed them to a mixture of university students, campus staff, and random volunteers they recruited.</p>
<p>And, at least on some level, the results indicate that women do like a larger penis. But the effect tails off quickly once the penis passes a given length, with the critical measurement being influenced by body height (the taller the person, the bigger the penis had to be). Complicating matters, women liked tall men, as well—it had just as strong an effect as penis size. So, that helped drag the preferred penis size to even greater lengths than it would have been otherwise.</p>
<p><strong>If you're emotionally attached to a Web service, using pot may be one of the least of your problems.</strong> A survey was used to check which of four factors was associated with college students growing emotionally attached to using Facebook. It came up with a simple answer: <a href="http://www.sciencedirect.com/science/article/pii/S0747563212003317">all of them</a>. Loneliness, anxiousness, alcohol, and marijuana use were all associated with getting attached to using the service, while loneliness and anxiousness were associated with establishing a lot of connections on it. Elsewhere in the same journal, we have a demonstration that <a href="http://www.sciencedirect.com/science/article/pii/S0747563211000379">Facebook users are narcissistic</a>. But they're also more extroverted and less lonely than non-users.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/weird-science-retracts-its-penis-to-keep-everyone-guessing/#p3">Read 4 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/weird-science-retracts-its-penis-to-keep-everyone-guessing/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=8jSlZ98435s:VBYLCdQZE3s:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=8jSlZ98435s:VBYLCdQZE3s:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=8jSlZ98435s:VBYLCdQZE3s:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=8jSlZ98435s:VBYLCdQZE3s:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=8jSlZ98435s:VBYLCdQZE3s:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=8jSlZ98435s:VBYLCdQZE3s:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/8jSlZ98435s" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/weird-science-retracts-its-penis-to-keep-everyone-guessing/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/weird-science-retracts-its-penis-to-keep-everyone-guessing/</feedburner:origLink></item>
		<item>
		<title>Robotics forecast: cool with a chance of lost humanity</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/zsfBe5zBuWg/</link>
		<comments>http://arstechnica.com/science/2013/04/robotics-forecast-cool-with-a-chance-of-lost-humanity/#comments</comments>
		<pubDate>Sat, 13 Apr 2013 16:10:44 +0000</pubDate>
		<dc:creator>Scott K. Johnson</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[Artificial Intelligence]]></category>
		<category><![CDATA[book review]]></category>
		<category><![CDATA[Robotics]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=243927</guid>
		<description><![CDATA[<i>Robot Futures</i> examines societal implications of revolutions in tech.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/robotfutures-640x938.jpg">
</div>
 <p>You might expect a book titled <i>Robot Futures</i> and written by a robotics researcher to be a whiz-bang prophecy of technologies that are the best thing since sliced bread. Soon we’ll be living to 200 while traveling from vacation to vacation in our flying cars. All the while, robots handle all the parts of our jobs that we hated anyway, right? Maybe, but this book isn't the place to find it. There’s plenty of speculation in it (I mean, we <i>are</i> talking about the future here) but it’s decidedly more pragmatic and sober than that.</p>
<p>The book’s author, <a href="http://www.cs.cmu.edu/~illah/">Illah Nourbakhsh</a>, runs Carnegie Mellon's Community Robotics, Education, and Technology Empowerment (<a href="http://www.cmucreatelab.org/">CREATE</a>) Lab, which “explores socially meaningful innovation and deployment of robotic technologies.” (He’s also one of the developers of the GigaPan imaging tools we <a href="http://arstechnica.com/science/2012/11/gigapans-gigapixel-images-bring-field-trips-to-geology-classrooms/">highlighted recently</a>.) The book, accompanied by Nourbakhsh’s <a href="http://www.robotfutures.org/">blog</a> of the same name, focuses on what he calls the “human side effects” of future technologies—the unexpected ways they could affect us socially and as a society.</p>
<p>Each chapter of the book—which progresses from the likely-seeming near future to an increasingly speculative and distant one— begins with a short, fictional story to illustrate the types of issues that could accompany a given technology. The rest of the chapter then beaks down those issues.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/robotics-forecast-cool-with-a-chance-of-lost-humanity/#p3">Read 10 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/robotics-forecast-cool-with-a-chance-of-lost-humanity/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zsfBe5zBuWg:E5AdQk0zC1A:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=zsfBe5zBuWg:E5AdQk0zC1A:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zsfBe5zBuWg:E5AdQk0zC1A:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=zsfBe5zBuWg:E5AdQk0zC1A:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zsfBe5zBuWg:E5AdQk0zC1A:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zsfBe5zBuWg:E5AdQk0zC1A:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/zsfBe5zBuWg" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/robotics-forecast-cool-with-a-chance-of-lost-humanity/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/robotics-forecast-cool-with-a-chance-of-lost-humanity/</feedburner:origLink></item>
		<item>
		<title>Diluting the scientific method: Ars looks at homeopathy (again)</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/zepdurVTUJw/</link>
		<comments>http://arstechnica.com/science/2013/04/the-pseudoscience-behind-homeopathy/#comments</comments>
		<pubDate>Sat, 13 Apr 2013 12:15:00 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Features]]></category>
		<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[homeopathy]]></category>
		<category><![CDATA[Pseudoscience]]></category>
		<category><![CDATA[science]]></category>

		<guid isPermaLink="false">http://arstechnica.com/science/2007/09/the-pseudoscience-behind-homeopathy/</guid>
		<description><![CDATA[In honor of World Homeopathy Week, we discuss why it shouldn't be celebrated.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap">
<div>
      <img src="http://cdn.arstechnica.net/wp-content/uploads/2013/04/recursive-water.jpg"><div class="caption" style="font-size:0.8em">
	
			<div class="caption-credit">
							Aurich Lawson				</div>
	  </div>
  </div>
 <p><em><p>Over the past couple of weeks, Ars has started receiving e-mails about a topic that doesn't really fit into our normal areas of coverage: homeopathy. The belief that the administration of nothing more than water can have a medicinal affect certainly isn't technology, and it only peripherally touches upon science, in the sense that any health benefits it provides seem to involve the placebo affect.</p>

<p>Nevertheless, many of the practitioners of homeopathy have tried to claim the mantle of science, creating scientific-sounding explanations for "water memory"—the idea that a substance can actually gain potency as it is diluted out of existence. These practitioners even formed a peer-reviewed journal to discuss their findings. Six years ago, the journal <em>Homeopathy</em> ran a special issue on these explanations, and the Ars writers teamed up to tackle the bizarre distortions of science that it contained.</p>

<p>So in honor of World Homeopathy Week (which we just received a press release for), we're going to run an updated version of that story. Not because we think there's a value in talking about homeopathy. Rather, by scientifically evaluating homeopathy's attempts to sound scientific and revealing it to be nothing more than "pseudoscience," we think it's possible to learn something about the scientific process and the reasoning that drives it. In turn, we can possibly learn to recognize other areas where scientific reasoning has ended up on the rocks.</p>

<p>This feature was originally published on September 11, 2007.</p></em></p>
<h2>Welcome to Waterworld</h2>
<p><span style="font-size: 14px; line-height: 19px;">Science can be a tricky thing to define, and it's sometimes easier to contrast it with some of the arguments that pose as science. Unfortunately, most of those issues are entangled with implications that keep the basic question—is this science?—obscured by emotional responses. Thus, the science of climatology has become entwined with political, economic, and policy issues. The science of evolution conflicts with the political and religious goals of some individuals. Even basic scientific questions about the nervous system get embroiled in family and personal health issues when topics like autism and radio frequency radiation are broached.</span></p>
<p>That's why a special edition of the journal <em>Homeopathy</em> appears to be a gift, allowing us to look at science and pseudoscience without getting entangled with politics and religion. Homeopathy claims to be a form of medical practice that's based on the principle that "like cures like." Given a set of symptoms, a homeopath will identify an herb or chemical that causes similar symptoms. Following a predefined ritual, the homeopath performs a series of dilutions of that chemical that continue well beyond the point where there should be no molecules of it left—the final solution is essentially well-shaken water.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/the-pseudoscience-behind-homeopathy/#p3">Read 75 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/the-pseudoscience-behind-homeopathy/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zepdurVTUJw:v1Of4Y5Tm2g:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=zepdurVTUJw:v1Of4Y5Tm2g:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zepdurVTUJw:v1Of4Y5Tm2g:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=zepdurVTUJw:v1Of4Y5Tm2g:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zepdurVTUJw:v1Of4Y5Tm2g:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=zepdurVTUJw:v1Of4Y5Tm2g:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/zepdurVTUJw" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/the-pseudoscience-behind-homeopathy/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/the-pseudoscience-behind-homeopathy/</feedburner:origLink></item>
		<item>
		<title>Want to name an exoplanet after your cat? IAU says dream on</title>
		<link>http://feeds.arstechnica.com/~r/arstechnica/science/~3/DwW5NDqXFy8/</link>
		<comments>http://arstechnica.com/science/2013/04/want-to-name-an-exoplanet-after-your-cat-iau-says-dream-on/#comments</comments>
		<pubDate>Fri, 12 Apr 2013 19:12:09 +0000</pubDate>
		<dc:creator>John Timmer</dc:creator>
				<category><![CDATA[Scientific Method]]></category>
		<category><![CDATA[astronomy]]></category>
		<category><![CDATA[exoplanets]]></category>

		<guid isPermaLink="false">http://arstechnica.com/?p=245103</guid>
		<description><![CDATA[Astronomy organization says paying to name exoplanets won't get you anything.]]></description>
				<content:encoded><![CDATA[<div id="rss-wrap"> <p>It has been incredibly exciting to bring you the latest news on the hunt for planets outside our Solar System. Over the past few years, the planets we've spotted have gotten closer to Earth's size, closer to the Solar System, and far, far more numerous. There's almost nothing that could dull the excitement—except, perhaps, the planets' names. Monikers like KOI-135 b and HAT-P-5 b don't exactly inspire a sense of the mysterious unknown.</p>
<p>Even the International Astronomical Union (IAU) admits they're a bit dull. "While exoplanet names such as 16 Cygni Bb or HD 41004 Ab may seem boring when considering the names of planets in our own Solar System, the vast number of objects in our Universe—galaxies, stars, and planets to name just a few—means that a clear and systematic system for naming these objects is vital," the organization has noted. But at least one company has decided that clear and systematic is not its thing, and it has started selling the chance to name a planet.</p>
<p>Or rather <a href="http://science.time.com/2013/03/07/name-your-own-exoplanet-for-4-99/">nominate a name</a> and download a certificate of said nomination. People can then vote on the nominees, with voting closing on Monday night. The company running the process says it donates some of its funds to astronomy research and education efforts.</p>
</div><p><a href="http://arstechnica.com/science/2013/04/want-to-name-an-exoplanet-after-your-cat-iau-says-dream-on/#p3">Read 1 remaining paragraphs</a> | <a href="http://arstechnica.com/science/2013/04/want-to-name-an-exoplanet-after-your-cat-iau-says-dream-on/?comments=1">Comments</a></p><div class="feedflare">
<a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=DwW5NDqXFy8:pCwlry6nOqU:V_sGLiPBpWU"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=DwW5NDqXFy8:pCwlry6nOqU:V_sGLiPBpWU" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=DwW5NDqXFy8:pCwlry6nOqU:F7zBnMyn0Lo"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?i=DwW5NDqXFy8:pCwlry6nOqU:F7zBnMyn0Lo" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=DwW5NDqXFy8:pCwlry6nOqU:qj6IDK7rITs"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=qj6IDK7rITs" border="0"></img></a> <a href="http://feeds.arstechnica.com/~ff/arstechnica/science?a=DwW5NDqXFy8:pCwlry6nOqU:yIl2AUoC8zA"><img src="http://feeds.feedburner.com/~ff/arstechnica/science?d=yIl2AUoC8zA" border="0"></img></a>
</div><img src="http://feeds.feedburner.com/~r/arstechnica/science/~4/DwW5NDqXFy8" height="1" width="1"/>]]></content:encoded>
			<wfw:commentRss>http://arstechnica.com/science/2013/04/want-to-name-an-exoplanet-after-your-cat-iau-says-dream-on/feed/</wfw:commentRss>
		<slash:comments>0</slash:comments>
		<feedburner:origLink>http://arstechnica.com/science/2013/04/want-to-name-an-exoplanet-after-your-cat-iau-says-dream-on/</feedburner:origLink></item>
	</channel>
</rss>
`