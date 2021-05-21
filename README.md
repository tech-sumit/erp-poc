# Jual.com - ERP
## ERP API GATEWAY

#### .env contents

    SERVICE_HOST=<SERVICE_HOST>
    SERVICE_PORT=<SERVICE_PORT>
    COOKIE_SECRET=<RANDOM_STRING>
    ACCOUNTS_HOST=<ACCOUNTS_SERVICE_HOST>

#### Execution Steps

For local server
    
    $ make start_local
For independant docker container
    
    $ make start_docker
For docker-compose stack

    $ make start_stack # Stack Build + Up
    $ make stop_stack # Stack Down
    $ make stack_logs # Stack Logs
