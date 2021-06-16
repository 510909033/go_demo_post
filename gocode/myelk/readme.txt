GET /test/_search
{
  "query": {
    "term": {
      "Name":"2*"
    }
  }
}

删除索引和数据
DELETE test
DELETE website


GET _search
{
    "query": {
        "query_string": {
            "query": "kill"
        }
    }
}

GET _search
{
    "query": {
        "query_string": {
            "query": "ford"
            , "fields": ["title"]
        }
    }
}

GET _search
{
    "query": {
      "term" : {
        "year" : 1962
    }
    }
}


GET _search
{
  "query": {
    "match_all": {}
  }
}

GET /website/_search
{
  "query": {
    "match_all": {}
  }
}

DELETE test
DELETE website

PUT /movies/movie/1
{
    "title": "The Godfather",
    "director": "Francis Ford Coppola",
    "year": 1972,
    "genres": ["Crime", "Drama"]
}
GET /movies/movie/9RAcpHQB_0iGhi__YlY0
GET /movies/movie/1

POST /movies/movie/6
{
    "title": "The Assassination of Jesse James by the Coward Robert Ford",
    "director": "Andrew Dominik",
    "year": 2007,
    "genres": ["Biography", "Crime", "Drama"]
}
DELETE /movies/movie/9RAcpHQB_0iGhi__YlY0
GET /movies/_search
GET _search
{
    "query": {
        "query_string": {
            "query": "drama"
        }
    }
}

GET _search
{
    "query": {
        "constant_score": {
            "filter": {
                "term": { "year": 1962 }
            }
        }
    }
}
GET _search
{
    "query": {
      "term" : {
        "year" : 1962
    }
    }
}

PUT /schools
POST /schools/_bulk
{"index": {"_index":"schools", "_type":"school", "_id":"1"}}
{"name":"Central School", "description":"CBSE Affiliation", "street":"Nagan","city":"paprola", "state":"HP", "zip":"176115", "location":[31.8955385, 76.8380405],"fees":2000, "tags":["Senior Secondary", "beautiful campus"],  "rating":"3.5"}

GET /schools/school/1

POST /schools/_bulk
{"index":{"_index":"schools", "_type":"school", "_id":"2"}}
{"name":"Saint Paul School", "description":"ICSE Afiliation", "street":"Dawarka", "city":"Delhi", "state":"Delhi", "zip":"110075","location":[28.5733056, 77.0122136], "fees":5000,"tags":["Good Faculty", "Great Sports"], "rating":"4.5"}
{"index":{"_index":"schools", "_type":"school", "_id":"3"}}
{"name":"Crescent School", "description":"State Board Affiliation", "street":"Tonk Road","city":"Jaipur", "state":"RJ", "zip":"176114","location":[26.8535922, 75.7923988],"fees":2500, "tags":["Well equipped labs"], "rating":"4.5"}







