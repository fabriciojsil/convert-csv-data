# convert-csv-data


 > How run test

``` shell
 make test
```

> Build project

``` shell
 make build
```

> Run built

``` shell
./csv-converter -file=path/to/file.csv -format=json -field=contact -order=asc
```

### Docker-compose environment

** At docker-compose.yml **

> **FILE** env to define csv file to be converted - *required

> **FORMAT** env to devine the output of file converted: accepeted values [json, xml] - json is the default value

> **ORDER** env to define the sorted output of file converted: accepeted values [asc, desc] - asc is the default value

> **FIELD** env to difine what field should be sorted

** Running at docker compose **

After setting the envs at Docker-compose run:

> docker-compose up --build
