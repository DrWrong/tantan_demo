# Solutions of the Tantan Back-End Devloper Test


## Db design

according to the requiremnt, I designed tow table the `user` table is used for record user information, and the `relationship` table is used for mantain relations between two users

Since I am a newcomer to the postgresql, (To be honst, I have just start to learn postgresql after I see the test) I won't use any specific feature of postgresql. But just treat is a commonly relationship database like mysql.

+ user table

    | field|  type | since| note|
	|------|-------|-------|------|
    |`id`|serial|`1.0.0`|the user id|
    |`name`|varchar(255)|`1.0.0`|the user name|


+ relationship table

    | field|  type | since| note|
	|------|-------|-------|------|
    |`id`|serial|`1.0.0`|auto increment id|
    |`uid1`|integer|`1.0.0`|link to the first user|
    |`uid2`|integer|`1.0.0`|link to the second user|
    |`state`|small int|`1.0.0`|the state of uid1 to uid2 1 stand for `disliked` 2 stand for `liked` 3 stand for `matched`|

    the uid1 together with uid2 should be a unique index


## Application Design

the application is desiged following the MVC pattern
In this application I used a framework called [`Monica`](https://github.com/DrWrong/monica)
This web framework is write by my own, aiming to make a productive yet simple WebFramework to deal with the daily RESTFUL stuff.
The idea of monica is based on a project called [`macaron`](https://github.com/go-macaron/macaron)

In this framework I realize the route, controller layer. As for orm layer, I borrowed [beego](http://beego.me/)

the package `ui` is the main logic of this application, in which there are two sub packagies called `contollers` and `models` for the controller layer and model layer

## useage

1. load db: in the proejct directory there is a file called `dbexport.pgsql` just import it
2. compile go code: cd to the project directory and issue `./install.sh`
3. make customized config: the config file locate in `conf/monica.yaml` open it and edit the `postgres` part and set the `dsn` according to your environment
4. just issue the `bin/ui.sh start` then server will run at port `8202` change the port to any you want in the config file


Thanks





