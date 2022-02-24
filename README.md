# Assignment 1: Context-sensitive University Search Service
By Daniel Hao Huynh

## Overview

In this assignment, I am going to develop a REST web application in Golang that provides the client to retrieve information about universities that may be candidates for application based on their name, alongside useful contextual information pertaining to the country it is situated in. For this purpose, I will interrogate existing web services and return the result in a given output format. 

The REST web services I have used for this purpose:
* http://universities.hipolabs.com/
  * Documentation/Source under: https://github.com/Hipo/university-domains-list/
* https://restcountries.com/
  * Documentation/Source under: https://gitlab.com/amatos/rest-countries

  My web service has three resource root paths: 

```
/unisearcher/v1/uniinfo/
/unisearcher/v1/neighbourunis/
/unisearcher/v1/diag/
```

Assuming the web service is hosted on localhost, on port 8080, my resource root paths would look something like this:

```
http://localhost:8080/unisearcher/v1/uniinfo/
http://localhost:8080/unisearcher/v1/neighbourunis/
http://localhost:8080/unisearcher/v1/diag/
````
