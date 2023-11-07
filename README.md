![SECURITY BY CURIOSITY](slides/img/security-by-curiosity.png)

# Why this repo?

The following repo contains instructions for a seminar on cybersecurity for the graduate program of 2023.


The goal is to feed your curiosity for hacking by studying red team tactics, a.k.a t. 

## From Curious to 1337-h4ck3r?

People often ask the internet "teach me how to hack" or "How do I become a hacker?". As graduates, you are curious by nature, you have what it takes. When you apply curiosity to use tech in unintended ways, then you are a hacker. When you make these applications practical for other people, thats when you have become 1337-hacker :)


# Basics

Below are some basics needed to better follow along the seminar on the 9th of November.

Happy learning

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
- [ ] [How Phishers Are Slinking Their Links Into LinkedIn](https://krebsonsecurity.com/2022/02/how-phishers-are-slinking-their-links-into-linkedin/)