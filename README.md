
# Introduction Red Team / Blue Team

The goal of this exercise is to make you more cyber security aware. To acheive this, you will do some red team exercises, the offensive side, and then we will do some blue team exercises, the defensive side. Security is a very large doctrine. 

Through your careers, people will try to sell security to you, either by some idea, which will cost your time, or a product which will cost both money and time. People will, with good intentions, misdirect your attention to perceived threats. Your duty as software engineers or IT practitioners is to keep yourselves well informed and to think critically when someone is selling you security, because [there is a lot of snake oil out there](https://en.wikipedia.org/wiki/Snake_oil)

The original scope of this exercise was limited to security by default. I do not know anything about that, so instead, I will teach you some of the things I have learned from my 3 years as a security analyst at one of the worlds largest managed security services providers.

For the rest of this exercise we will keep it simple and simulate a cyber attack : )

Happy hacking

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

# Attack Scenario: Zeegma Corporate espionage against Knowit Solutions Cocreate

You are cyber mercenaries who have been contracted to perform corporate espionage against Knowit Solutions Cocreate. Your employer, Zeegma, has been falling behind its competitor Cocreate and is desperate to know why all its customers have left to Knowit Solutions Cocreate. Your **primary objective** is to compromise Cocreate's customer relationship management platform (CRM) and emails for individuals who work with the CRM system. Using the compromised accounts, you will relay customer relations information to Zeegma, and brief them on your findings.

The **secondary objective** is to establish persistance inside Cocreate's environment so that the customer Zeegma can repeatedly extract valuable business intelligence from the victom.

## Action Plan

Here are some steps that he cyber mercenaries would make to acheive the goal:

- *Recon* Find some email addresses of Cocreate
- Send targeted phishing emails
- *Initial Foothold*: Login as employees using compromised 2fa tokens
- **Primary Objective**: Search emails, teams and other instant messages for references to the CRM system.
- **Secondary Objective**, *Persistance*: Send PMs from compromised accounts with attached malware (abuse trust relationships)
- Log keystrokes to extract more credentials
- **Primary Objective**: Extract cookies from browser to login to CRM using malware

## Get Knowit Emails

To send phishing emails, we need target addresses. There are multiple good ways of finding out addresses of the target organisation. Below is an example:

### Go to [this link](https://phonebook.cz/)

Use these credentials:

`simplelogin.74q6u@simplelogin.com:8n6LYeZecRbcAD`

### Search for email/domains/urls with domain `knowit.se`

You get a long list of email addresses. What extra information do you need to add to this list to make it good enough for our spearphishing attack towards Cocreate?

### Info to craft good spear phishing emails

To design a good phishing email, attackers can use publicly known knowledge about individials or about the company. Typically, you would research trust relationships and hierarchies and spoof emails from a manager to subordinates.

- where would you go on the internet to find sensitive info to craft spearphishing emails?


## Initial Foothold & Persistance

This part of the exercise will be demoed using our own emails during the presentation on the 9th of November 2023.

# Extra: Exploiting Vulnerable Services

Compromising organisations through phishing or other social engineering means is probably the easiest and most effective method of establishing a foothold in an organisation. However, with the advent of modular exploit delivery frameworks like metasploit, finding vulnerable targets on the internet is very easy and has lowered the barrier of entry to mass-exploitation for threat actors. Anyone can do this technique with access to some rotating proxies and a credit card for cloud services.

## Attack Plan

- Compile a list of common vulnerable services
- Write a scanner that scans the entire Ipv4 range 
- Scan the entire internet for the vulnerable services, rotating proxies as the your ip addresses get blacklisted
- Compile dataset on exploited machines on the internet
- Research each exploited service, who does it belong to, how much money do they make
- Sell the initial access to other threat actors on hacker forums / Dark web


# Recommended Reading

- [ ] [Open-Source Intelligence (OSINT) in 5 Hours - Full Course - Learn OSINT! ](https://www.youtube.com/watch?v=qwA6MmbeGNo&t=7429s)
- [ ] [Hack the box, pluralsight for hacking](https://www.hackthebox.com/)