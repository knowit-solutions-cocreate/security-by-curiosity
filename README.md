![SECURITY BY CURIOSITY](slides/img/security-by-curiosity.png)

# Why this repo?

The following repo contains instructions for a seminar on cybersecurity for the graduate program of 2023.


The goal is to feed your curiosity for hacking by studying red team tactics, a.k.a the hacker/evil side 👹. 

## Why is curiosity good for security?

When you apply curiosity to use tech in unintended ways, i.e hacking, then you are internalizing knowledge about systems and how they function. This is arguably the best way to learn about cyber security.

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

The **secondary objective** is to establish persistance inside Cocreate's environment so that the customer Zeegma can repeatedly extract valuable business intelligence from the victim.

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


### Task: Write your own phishing email

Your task is to use all knowledge you have about your current colleagues and write a phishing email that maximises click through rate. Bring it to the seminar on 9th and we will read through them together.


## Initial Foothold & Persistance

This part of the exercise will be demoed using our own emails during the presentation on the 9th of November 2023.


# Recommended Reading
- [ ] [The Jargon File](http://www.catb.org/jargon/html/index.html)
- [ ] [Open-Source Intelligence (OSINT) in 5 Hours - Full Course - Learn OSINT! ](https://www.youtube.com/watch?v=qwA6MmbeGNo&t=7429s)
- [ ] [Hack the box, pluralsight for hacking](https://www.hackthebox.com/)
- [ ] [How Phishers Are Slinking Their Links Into LinkedIn](https://krebsonsecurity.com/2022/02/how-phishers-are-slinking-their-links-into-linkedin/)
- [ ] [How to set up Evilginx to phish Office 365 credentials](https://janbakker.tech/how-to-set-up-evilginx-to-phish-office-365-credentials/)
