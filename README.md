# 307: Temporary Redirect

One fine evening while looking up response codes, I thought buying a domain for one of the short phrases associated with the response codes would be *neat*…

This repo hosts the content for `temporaryredirect.com` and the location it redirects to (`307.temporaryredirect.com`) using a `307` response. So meta.

```sh
> curl -i https://temporaryredirect.com
HTTP/1.1 307 Temporary Redirect
Location: https://307.temporaryredirect.com
Server: cloudflare-nginx

… other headers

```

![307 Temporary Redirect Website Screenshot](misc/website-safari-screenshot.png)

The page at `307.temporaryredirect.com` is static content in the `gh-pages` branch of this repo and is served by Github Pages. The `307` response from `temporaryredirect.com` is returned by the Heroku *app* at `temporary-redirect.herokuapp.com`, code for which is also in this repo. Heroku is setup to auto-deploy changes to this repo.

The really powerful consequence of this setup is that changes can be made using Github’s in-browser code editor and on commit, changes are deployed automatically to both the static site and the tiny app on Heroku. Super useful for minor fixes and small feature additions. No deployments, FTP uploads (remember those?), git pushes to deal with. Just edit and save (with a descriptive commit message of course). Done.

## Credits
* Github (free public repo hosting, Github Pages)
* Heroku (free app hosting*)
* CloudFlare (free SSL!, simple DNS record management, free CNAME flattening)
* Skeleton CSS (for a simple, light weight, responsive CSS *framework*)
* Mozilla (for the Fira Mono font)
* Google Fonts (for free web font hosting)

*Heroku has a free tier but with limitations – they *turn off / pause* the dyno (container?) that the app is running on when there is no activity. In addition they have a per day limit on the number of hours it can be active. This in practice leads to a several second delay in response times for websites that see very little usage when the dyno needs to be spun back up. In spite of this limitation, the toolchain that surrounds Heroku’s runtime is very powerful and makes them a compelling choice*