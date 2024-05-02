# 🌄 Background

CatsSocial adalah aplikasi dimana pemilik kucing dapat saling menjodoh-jodohkan kucingnya

---

# 📝 Requirement

Berikut adalah fungsi-fungsi yang diharapkan oleh user kita

## 🫡 Functional

Berikut adalah fungsi-fungsi untuk kebutuhan produk

### Authentication & Authorization

**Background:**

Karena mayoritas pemilik kucing seringkali hanya menjodohkan setahun 2x, jadi metode login email dipilih sebagai cara pendaftaran

**Contract:**

- POST /v1/user/register
    - Request

        ```jsx
        {
        	"email": "email", // not null, can't be duplicate email, should be in email format
        	"name": "namadepan namabelakang", // not null, minLength 5, maxLength 50, name can be duplicate with others
        	"password": "" // not null, minLength 5, maxLength 15
        }
        ```

    - Response
        - `201` User successfully registered

        ```jsx
        {
            "message": "User registered successfully"
            "data": {
        			"email": "email@email.com", 
        			"name": "namadepan namabelakang", 
              "accessToken": "qwertyuiopasdfghjklzxcvbnm" // token should have 8 hours until expiration
            }
        }
        
        ```

    - Errors:
        - `409` conflict if email exists
        - `400` request doesn’t pass validation
        - `500` if server error
- POST /v1/user/login
    - Request

    ```jsx
    {
    	"email": "email", // not null, can't be duplicate email, should be in email format
    	"password": "" // not null, minLength 5, maxLength 15
    }
    ```

    - Response
        - `200` User successfully logged

        ```jsx
        {
            "message": "User logged successfully"
            "data": {
        			"email": "email@email.com",
        			"name": "namadepan namabelakang", 
              "accessToken": "qwertyuiopasdfghjklzxcvbnm" // token should have 8 hours until expiration
            }
        }
        ```

    - Error
        - `404` if user not found
        - `400` if password is wrong
        - `400` request doesn’t pass validation
        - `500` if server error

---

### Manage Cats

<aside>
💡 All request here should use `Bearer Token` from `accessToken` auth route

</aside>

Background:

Setelah mendaftar, user dapat mendaftarkan kucingnya untuk dicarikan jodohnya

Contract:

- POST /v1/cat
    - Request Body

    ```jsx
    {
    	"name": "", // not null, minLength 1, maxLength 30
    	"race": "", /** not null, enum of:
    			- "Persian"
    			- "Maine Coon"
    			- "Siamese"
    			- "Ragdoll"
    			- "Bengal"
    			- "Sphynx"
    			- "British Shorthair"
    			- "Abyssinian"
    			- "Scottish Fold"
    			- "Birman" */
    	"sex": "", // not null, enum of: "male" / "female"
    	"ageInMonth": 1, // not null, min: 1, max: 120082
    	"description":"" // not null, minLength 1, maxLength 200
    	"imageUrls":[ // not null, minItems: 1, items: not null, should be url
    		"","",""
    	]
    }
    ```

    - Response
        - `201` successfully add cat

            ```jsx
            {
            	"message": "success",
            	"data": 
            		{
            			"id": "", // use whatever id
            			"createdAt": "" // should in ISO 8601 format
            		}
            }
            ```

        - `400` request doesn’t pass validation
        - `401` request token is missing or expired
- GET /v1/cat
    - Request
        - Param (all optional)
            - `id` limit the output based on the cat’s id
                - value should be a string
            - `limit` & `offset` limit the output of the data
                - default `limit=5&offset=0`
                - value should be a number
            - `race` filter based on race
                - enum of `"Persian"`|`"Maine Coon"`|`"Siamese"`|`"Ragdoll"`|`"Bengal"`|`"Sphynx"`|`"British Shorthair"`|`"Abyssinian"`|`"Scottish Fold"`|`"Birman"`
            - `sex` filter based on sex
                - enum of `"male"` | `"female"`
            - `hasMatched` filter based on match
                - enum of `true`|`false`
            - `ageInMonth` filter based on age
                - value should be a string, but parsed like this:
                    - `ageInMonth=>4` searches data that have more than 4 months
                    - `ageInMonth=<4` searches data that have less than 4 months
                    - `ageInMonth=4` searches data that have exact 4 month
            - `owned` filter based on the cat that the user own
                - value should be `true`/`false`
            - `search` display information that contains the name of search
                - value should be a string
    - Response
        - `200` successfully get cats

            ```jsx
            {
            	"message": "success",
            	"data": [ // ordered by newest first
            		{
            			"id": "", // use whatever id
            			"name": "", 
            			"race": "", 
            			"sex": "", 
            			"ageInMonth": 1, 
            			"imageUrls":[ 
            				"","",""
            			],
            			"description":"" // not null, minLength 1, maxLength 200
            			"hasMatched": true, // true if the cat is already matched
            			"createdAt": "" // should in ISO 8601 format
            		}
            	]
            }
            ```

        - `401` request token is missing or expired
- PUT /v1/cat/{id}
    - Request Path Params
        - `id` the id that user want to edit
    - Request Body

    ```jsx
    {
    	"name": "", // not null, minLength 1, maxLength 30
    	"race": "", /** not null, enum of:
    			- "Persian"
    			- "Maine Coon"
    			- "Siamese"
    			- "Ragdoll"
    			- "Bengal"
    			- "Sphynx"
    			- "British Shorthair"
    			- "Abyssinian"
    			- "Scottish Fold"
    			- "Birman" */
    	"sex": "", // not null, enum of: "male" / "female"
    	"ageInMonth": 1, // not null, min: 1, max: 120082
    	"description":"" // not null, minLength 1, maxLength 200
    	"imageUrls":[ // not null, minItems: 1, items: not null, should be url
    		"","",""
    	]
    }
    ```

    - Response
        - `200` successfully add cat
        - `400` request doesn’t pass validation
        - `401` request token is missing or expired
        - `404` id is not found
        - `400` sex is edited when cat is already requested to match
- DELETE /v1/cat/{id}
    - Request Path Params
        - `id` the id that user want to delete
    - Response
        - `200` successfully delete cat
        - `401` request token is missing or expired
        - `404` id is not found

---

### Match cat

<aside>
💡 All request here should use `Bearer Token` from `accessToken` auth route

</aside>

**Background:**

Setelah user mendapatkan kucing yang mau dijodohkan, user dapat melakukan request untuk menjodohkan kucingnya dengan melakukan janjian

**Contract:**

- POST /v1/cat/match
    - Request

    ```jsx
    {
    	"matchCatId": "",
    	"userCatId": "",
    	"message": "" // not null, minLength: 5, maxLength: 120
    }
    ```

    - Response
        - `201` successfully send match request
        - `404` if neither `matchCatId` / `userCatId` is not found
        - `404` if `userCatId` is not belong to the user
        - `400` if the cat’s gender is same
        - `400` if either `matchCatId` / `userCatId` already matched
        - `400` if `matchCatId` / `userCatId` is from the same owner
        - `401` request token is missing or expired
- GET /v1/cat/match
    - Response body
        - `200` successfully get match requests

            <aside>
            ⚠️ This should show in both the issuer and the receiver

            </aside>

            ```json
            {
            	"message": "success",
            	"data": [ // ordered by newest first
            		{
            			"id": "", // use whatever id
            			"issuedBy": {
            				"name": "",
            				"email": "",
            				"createdAt": "" // should in ISO 8601 format
            			},
            			"matchCatDetail": {
            				"id": "",
            				"name": "", 
            				"race": "", 
            				"sex": "", 
            				"description":"",
            				"ageInMonth": 1, 
            				"imageUrls":[ 
            					"","",""
            				],
            				"hasMatched": false,
            				"createdAt": "" // should in ISO 8601 format
            			},
            			"userCatDetail": {
            				"id": "",
            				"name": "", 
            				"race": "", 
            				"sex": "", 
            				"description":"",
            				"ageInMonth": 1, 
            				"imageUrls": [ 
            					"","",""
            				],
            				"hasMatched": false,
            				"createdAt": "" // should in ISO 8601 format
            			},
            			"message": "",
            			"createdAt": "" // should in ISO 8601 format
            		}
            	]
            }
            ```

        - `401` request token is missing or expired
- POST /v1/cat/match/approve
    - Request Body

        <aside>
        ⚠️ Once a match is approved, other match request that matches both the issuer and the receiver cat’s, will get removed

        </aside>

        ```json
        {
        	"matchId":""
        }
        ```

    - Response
        - `200` successfully matches the cat match request
        - `400` `matchId` is no longer valid
        - `401` request token is missing or expired
        - `404` `matchId` is not found
- POST /v1/cat/match/reject
    - Request Body

        ```json
        {
        	"matchId":""
        }
        ```

    - Response
        - `200` successfully reject the cat match request
        - `400` `matchId` is no longer valid
        - `401` request token is missing or expired
        - `404` `matchId` is not found
- DELETE /v1/cat/match/{id}

    <aside>
    ⚠️ Match can only be deleted by issuer

    </aside>

    - Request Path Params
        - `id` the id that user want to remove
    - Response
        - `200` successfully remove a cat match request
        - `400` `matchId` is already approved / reject
        - `401` request token is missing or expired
        - `404` `matchId` is not found

---

## 👾 Non Functional

Berikut adalah fungsi teknis yang perlu dilakukan

### Backend

- Backend server should:
    - Use `Golang` (use any web framework that you want)
    - Use `Postgres` database
    - Run in port `8080`
    - Not using ORM / Query generator, only raw query
    - No external caching (redis/memcached)
- Use these env variable for production server compatibility:

    ```bash
    export DB_NAME=
    export DB_PORT=
    export DB_HOST=
    export DB_USERNAME=
    export DB_PASSWORD=
    export DB_PARAMS="sslmode=disabled" # this is needed because in production, we use `sslrootcert=rds-ca-rsa2048-g1.pem` and `sslmode=verify-full` flag to connect
    # read more: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.Concepts.General.SSL.html
    export JWT_SECRET=
    export BCRYPT_SALT=8 # don't use 8 in prod! use > 10
    
    ```

- After you done, compile it with the binary name “`main`" by env `GOARCH=amd64` and `GOOS=linux`
    - Then, send the binary to the server by `scp` (a temporary server will be available on Friday for testing)
- Use any folder structure that you want
- We will use RDS Postgres database by `db.t2.medium` spec
- We will use EC2 Ubuntu server by `t2.medium` spec

---

### Database Migration

Database migration must use [golang-migrate](https://github.com/golang-migrate/migrate) as a tool to manage database migration

- **Short Tutorial:**
    - Direct your terminal to your project folder first
    - Initiate folder

        ```bash
        mkdir db/migrations
        
        ```

    - Create migration

        ```bash
        migrate create -ext sql -dir db/migrations add_user_table
        
        ```

      This command will create two new files named `add_user_table.up.sql` and `add_user_table.down.sql` inside the `db/migrations` folder

        - `.up.sql` can be filled with database queries to create / delete / change the table
        - `.down.sql` can be filled with database queries to perform a `rollback` or return to the state before the table from `.up.sql` was created
    - Execute migration

        ```bash
        migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
        
        ```

    - Rollback migration

        ```bash
        migrate -database "postgres://username:password@host:port/dbname?sslmode=disable" -path db/migrations up
        
        ```


<aside>
⚠️ At showcase day, you should migrate the production database yourself in your local environment
And after you done, you should also rollback the migration yourself

</aside>