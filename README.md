# Twitter Searcher

This project demonstartes the usage of Insights for Twitter service on [Bluemix] using Golang

## Usage

1. Create a new CloudFoundary Go web app on Bluemix
2. Download the starter code.
3. Copy `manifest.yml` from the starter code package.
4. Past it in the root folder of the project.
5. Add Insights for Twitter service to the app
6. Copy the services URL
7. Add the following `(/api/v1/messages/search?q=)` to the end of the URL. It will look like `https://XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXX:XXXXXXXXXX@cdeservice.mybluemix.net/api/v1/messages/search?q=`
8. Past the final URL in `DEF_URL` constant.
9. Push the application to the app using CloudFoundary CLI `cf push`

Start searching!

[Bluemix]: https://bluemix.net

