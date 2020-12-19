# registry-auth

A no-op token auth server mod based on https://github.com/adigunhammedolalekan/registry-auth for testing purpose.

NOTE the token auth server:
1. is supposed to be deployed with a TLS-enabled registry via docker compose;
2. shares the same TLS cert and private key with the registry;
3. does no authentication and pass-through authorization.

At the minimal level one can employ self-signed cert and the corresponding private key for token auth to function. Detailed how-to see `docker-compose.yml`

(The token auth mechanism employed by Docker registry is based on [Json Web Token](https://jwt.io/introduction))
