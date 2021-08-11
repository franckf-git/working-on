With the file enwiki-20210620-abstract.xml (6Go) from the autosave of [Wikipedia](https://meta.wikimedia.org/wiki/Data_dumps) , we will parse it and save datas in a sqlite database and export stats in data visualization.
It is a training on larger data in go code.

# parsing

## source file

xml file start with a `feed` anchor them list terms with the tags:

```
<doc>
<title>Wikipedia: Mariapia Degli Esposti</title>
<url>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti</url>
<abstract>Cancer Council of Western Australia Cancer Researcher of the Year 2017Research Excellence Awards - Cancer Council Western Australia (2017)</abstract>
<links>
<sublink linktype="nav"><anchor>Research</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Research</link></sublink>
<sublink linktype="nav"><anchor>Education</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Education</link></sublink>
<sublink linktype="nav"><anchor>Occupation</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Occupation</link></sublink>
<sublink linktype="nav"><anchor>Current</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Current</link></sublink>
<sublink linktype="nav"><anchor>Former</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Former</link></sublink>
<sublink linktype="nav"><anchor>Personal life</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#Personal_life</link></sublink>
<sublink linktype="nav"><anchor>References</anchor><link>https://en.wikipedia.org/wiki/Mariapia_Degli_Esposti#References</link></sublink>
</links>
</doc>
```

`title`, `url` and `abstract` will be easy to parse and put in database. `links` and the `sublinks` are more trickies (not really but put them on the side for now) we will just save the count of `sublinks` in each `doc` for now.

## database structure

From the structure of the xml, the database structure look like:
table `doc`

```
id int
title string
url string
abstract string
links int
```

table `unknown`

```
id int
unknowntag string
iddoc int // relationnal with `doc`
```

> database, open and insert DONE

## others

- What about unknown tags ? > they will be ignore for the original parsing but they will be store in another table.
- What about if the source file is corrumpt, it is an autosave of wikipedia, like missing closing tags ? > the errors will be handle in the code.
- TDD ? > a lot of new things for eme (parsing xml, sqlite, data visualization), so no, not this time.
- concurrency ? > it will be a perfect case, but start single, if the performances are bad we will come back to it (and I am not sure sqlite like multi writing).
