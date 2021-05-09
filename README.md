# Vagon Backend

Powered by Golang, Cloud Run on GCP and CockroachDB

# Database

Database connection is with a string using the environment variable `DB_CONN`

You can set this variable when running the application as a docker cointaner after you `make build` OR with a .env file in a folder named secrets

We 💗 CRL and CRDB (especially Amruta)

# 🔗Important Links 🔗

Github Org: https://github.com/TOHacks-Team-Alpha

Deployed instance: https://vagon.tech

API: https://api.vagon.tech/ping

Video: https://youtu.be/tBdPE6mJ51w

Slides: https://docs.google.com/presentation/d/1SLyeZgzB3a9gklk52bidCC0dDGhsr9-jrtO2YBWGbKs

# ✨ Inspiration ✨

We were inspired by the a problem that millions across the world face each day. Every morning and evening, people take their cars alone each day, while millions are forced to spend time commuting by public transportation. Not only does this create issues in equity, but it has come to seriously harm our environment through greenhouse gasses. Vagon (derived from the Armenian translation for 'car') provides a paradigm shift in transportation by empowering both car owners and public transportation users to find ways to carpool _ without _ needing to change one's lifestyle.

# ⚒ What it does ⚒

Vagon allows all users to be both a **rider** and a **driver**. Users can create **drives** which **riders** can request to join. When looking for a rider, users provide their location and a maximum distance to which they are willing to travel to the departure address and how far they are willing to be dropped off to their destinations address.

Both drivers and riders are incentivized in this system using our EcoCoin. Users can redeem their EcoCoin for gift cards, tree planting, Vagon swag or donations. Through this system, all users are incentivized to use Vagon rather than riders paying on traditional ride sharing platforms.

# 👷‍♂️ How we built it 👷‍♀️

The platform was built using Vue.js on the Nuxt.js framework on the front-end. This system makes it easy to rapidly develop an app with data visualization tools such as charts and maps (thanks to the Vue Google Components), without compromise on speed or SEO ability (thanks the the server side rendering).

The backend for logic was built in Golang with help from the Gin framework. Our Golang JSON REST API communicates to the frontend to send and receive all user information. In total, we developed a dozen end-points that require JWT authentication.

User accounts are created and managed by **Google Firebase 🔥**, allowing for out-of-the-box verification and account resetting. As well, firebase serves to authenticate user requests by providing the frontend with a token upon log-in that can be passed to the Golang backend and the once again verified against the firebase servers.

User data is saved with **CockroachDB hosted on Google Cloud** to allow a fast, yet secure, reliable and resilient data store. This is useful to people who may access the web app from a variety of devices. As well, the information can be downloaded by the user to perform their own analysis. As the platform is designed to be used around the word, CockroachDB can be leveraged to provide results quickly based on people's location.

The last part of the application is deployment which is handled by using **Docker and Google Cloud Run**. A single GCR service supports up to 100 000 users at maximum scale, thus our app can already be used by thousands of people across the world.

# 💪 Challenges we ran into 💪

For this hackathon, we took on a project with a complicated user flow that requires back and forth between users and drivers. Developing this meant consistently iterating the design to stay user focused leading to changes in the frontend, backend and database schema. Overall, its never easy to develop a whole app in 24 hours, but we learned a lot about rapid development and teamwork.

# 🥇 Accomplishments that we're proud of 🥇

We are super proud that the project is deployed and accessible for anyone to use! As well, we were able to apply our knowledge of SWE that we learned in university through using tools to test the accessibility of the web site and properly access use cases.

# 🚸 What we learned 🚸

Working with so much to do taught us a lot about project management. We set lofty goals and although it was rewarding to see progress made, the overall lack of "sprints" caused us to waste time and go back to redo frontend, backend and database components.

# 😲 What's next for Vagon 😲

Creating a dual sided marketplace for equitable transportation is a lofty goal. To have a project like this gain traction, we would need to partner with the government and NGOs. As well, there are other features such as asking drivers to detour for pick-up and drop-off.
