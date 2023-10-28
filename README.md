
# Introduction Red Team / Blue Team

The goal of this exercise is to make you more cyber security aware. To acheive this, you will do some red team exercises, the offensive side, and then we will do some blue team exercises, the defensive side. Security is a very large doctrine. 

Through your careers, people will try to sell security to you, either by some idea, which will cost your time, or a product which will cost both money and time. People will, with good intentions, misdirect your attention to perceived threats. Your duty as software engineers or IT practitioners is to keep yourselves well informed and to think critically when someone is selling you security, because [there is a lot of snake oil out there](https://en.wikipedia.org/wiki/Snake_oil)

The original scope of this exercise was limited to security by default. I do not know anything about that, so instead, I will teach you some of the things I have learned from my 3 years as a security analyst at one of the worlds largest managed security services providers.

We will simulate a cyber attack against fictitious and real targets within our organization Knowit Solutions Cocreate. Don't worry, We have gotten permission from the Knowit Cheif Security Officer.


# Instructions

- [ ] Download this repo
- [ ] Make a branch with your name

You can now start working on the exercises.

This guide will make you good at security. Joking. But it will give you a good foundation. It will probably make you more security aware than moste people around you though.

# Basics
 
Learn the following:

- [ ] What is dns?
- [ ] What is localhost?
- [ ] What is RFC1918?
- [ ] What are ports?
- [ ] What are the most common services you find on remote systems? What ports do they run on?
- [ ] What is the default shell on Ubuntu?
- [ ] What is the default shell on Microsoft?
- [ ] What is priveledge escalation?
- [ ] What is remote code execution?
- [ ] [What is the cyber kill chain?](https://www.eccouncil.org/cybersecurity-exchange/threat-intelligence/cyber-kill-chain-seven-steps-cyberattack/)


# Attack Scenario: Zeegma Corporate espionage against Knowit Solutions Cocreate

You are cyber mercenaries who have been contracted to perform corporate espionage against Knowit Solutions Cocreate. Your employer, Zeegma, has been falling behind its competitor Cocreate and is desperate to know why all its customers have left to Knowit Solutions Cocreate. Your objectives are the following:

- Establish an Initial foothold in Cocreate's environment
- Compromise Cocreate's CRM system
- Compromise Cocreate's internal IT systems
- Achieve persistance, either backdooring the user's workstations or internal IT systems
- Report back to Zeegma about with your findings  



- Reconnaiscance, information about the target is collected for analysis
- 
## Initial Foothold: Recon & Osint 
## 


## Recon & Pivoting through a victim network

- [ ] Watch this video, [How To Pivot Through a Network with Chisel](https://www.youtube.com/watch?v=pbR_BNSOaMk)
- [ ] Do a ping sweep on the Knowit corporate network, what hosts are up? Can you find other graduates machines?
- [ ] Can you find the Invativa Network Attached Storage (NAS) host using nmap? (optional) use `nmap -p445 <network>` for SMB scanning
- [ ] Write down your findings

> [!IMPORTANT] 
> Be careful when nmaping. Nmap will trigger alerts for network intrusion detection systems.

## Hacking basics & Hack the box

After the reconnaiscance phase an attackaer

There is no feeling better than poping a shell on a remote system, [Hack The Box](https://www.hackthebox.com) is a place where you can do that without getting arrested. Here are two boxes I recommend doing

- [ ] Blue, ms17-010, WannaCry Exploit, Here's a walkthrough [Ippsec walkthrough](https://www.youtube.com/watch?v=YRsfX6DW10E)
- [ ] Shocker, Shellshock, Heres a walkthrough [Ippsec walkthrough](https://www.youtube.com/watch?v=IBlTdguhgfY)

The first box demonstrates how trivial it is to pwn a system with one of the NSA tools Bluekeep which was leaked in 2017. The second box demonstrates the shellshock exploit, which people are still scanning for to this day, even though the exploit is almost 10 years old.

>[!NOTE]
> The Ippsec guides can be a bit thick for a beginner. Just do an online search for a walkthrough. Do not try to hack on your own if you do not know what you are doing.

# Open Source Intelligence gathering (OSINT) Course

TODO Cyber mentor OSINT Guide, 

# Mapping Knowit Infrastructure

# Phishing & Spear Phishing

EvilNginx

# Getting Knowit Email Addresses

- [ ] Get the list of emails from open source tools
- [ ] Who is working in HR / Office?
- [ ] Who is in management?
- [ ] Who are the devs?
- [ ] Use methods from OSINT guide
- [ ] Brute force the Fake Webserver for access

# Anti-virus evasion and hooking OS Api calls (advanced topic)
