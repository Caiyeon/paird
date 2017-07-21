### Introduction

Paird was a proof of concept slack application built in 2017 SAP hackathon. This was built within 24 hours.

It uses slack as a platform to facilitate pairing people up for mentorship purposes.

---

### Problem

Mentorship programs within organizations take an enormous amount of HR resources. A pairing of 200 participants can easily take dozens or hundreds of hours.

In addition, mentors and mentees may feel less inclined to join because of the work required to signup.

Also, emails are so 2016.

### Solution

Paird is a slack app, and people can interact with it on slash commands.

People can signup at their own leisure, as mentor/mentee, looking for a mentor/mentee

People can add tags to themselves, and search for specific tags. They can also list all current popular tags within the organization to see what is in demand

Dietary and location options are also available. When a pairing is made by the backend server, if location or dietary options are set, suggestions of nearby places will be presented to the pairing as potential meetups.

---

### Tech stack

* **Storage**: Paird uses boltdb (key & value storage), because SQL databases are also very 2016.
* **Frontend**: VueJS is used for the frontend, with vue-bootstrap template
* **Backend**: Golang is used for the backend
* **Host**: Previously hosted on DigitalOcean, currently hosted on GitHub pages
    - obviously the API server will no longer work, but it's free!
* **Slack**: Slack bot API is used for slash commands and interactive messaging
* **Yelp**: Yelp API is used to get a list of nearby restaurants
* **Vault**: Vault is used for encryption as a service, because why not?
* **CircleCI**: Used for continuous integration and also automatically built binaries

Special thanks to Mozilla for certificates from Let's Encrypt

And credits to Yonni Luu for the following diagram:
![](screenshots/tech_stack.png)
