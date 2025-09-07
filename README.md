# TON OAuth Service

![TON OAuth](https://img.shields.io/badge/TON-OAuth-blue)

**TON OAuth Service** is a project that enables secure user authentication through **TON wallets**. The project implements an OAuth-like authorization approach, using TON wallet signatures instead of passwords.

## ⚙️ Installation

1. **Clone the repository**
```bash
git clone https://github.com/dadilll/TONKeyAuth.git
cd TONKeyAuth
```
2. **Generate RSA keys**
```bash
mkdir -p key
openssl genrsa -out keys/private.pem 2048
openssl rsa -in keys/private.pem -pubout -out keys/public.pem
```
3 **Running the Service**
```bash
docker-compose build
docker-compose up -d
```

## ⚙️ Configuration

The service uses a configuration file to set up its parameters. The main settings include:

- **PrivateKeyPath** – path to the RSA private key used for signing JWT tokens (e.g., `key/private.pem`).  
- **PublicKeyPath** – path to the RSA public key used for verifying JWT tokens (e.g., `key/public.pem`).  
- **HTTPServerPort** – port on which the service listens (e.g., `8080`).  
- **Issuer** – the issuer name included in generated JWT tokens (e.g., `TON-OAUTH`).  
- **KeyName** – name of the key used in JWKS responses (e.g., `main-key`).  

## 📄 API Documentation 

The full API documentation for **TON OAuth Service** is available through **Swagger UI**.  
After running the service (see Installation section), open the following URL in your browser:
This interface allows you to explore all endpoints, view request/response schemas, examples, and test the API directly.

---

### 🔹 Endpoints Overview

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/oauth/authorize` | GET | Initiate authorization and generate a one-time challenge (nonce) for the user. |
| `/oauth/verify` | POST | Verify signed message from TON wallet using ed25519 signature. |
| `/oauth/token` | POST | Issue a JWT token after successful verification of a TON wallet. |
| `/oauth/jwks` | GET | Retrieve JSON Web Key Set (JWKS) containing public keys for JWT verification. |
| `/oauth/verify-token` | POST | Verify a JWT token issued by the service. |

 Each endpoint includes detailed request/response examples and validation rules, which are available in the Swagger UI.

## 🚀 Usage / How it works

**TON OAuth Service** allows users to securely authenticate using their TON wallets without traditional registration. The service works as follows:

1. **Authorization initiation**  
   The frontend or client requests authorization from the service, providing a redirect URI for post-login.

2. **Challenge generation**  
   The backend generates a unique one-time challenge (nonce) for the user.

3. **Signing the challenge**  
   The user signs the challenge using their TON wallet. This can be done via TonConnect SDK on the frontend.

4. **Sending the signature to the backend**  
   The signed challenge and the user's public key are sent to the backend for verification.

5. **Signature verification**  
   The backend validates the signature against the provided public key. Optionally, it can check whether the wallet exists on the TON blockchain.

6. **JWT issuance**  
   After successful verification, the backend issues a JWT signed with its private key, which the client can use for authenticated requests.

7. **Secure interaction**  
   Tokens have a limited lifetime, are signed, and cannot be tampered with, providing security against replay attacks and unauthorized access.

## 🔒 Security Considerations

**TON OAuth Service** is designed with security and privacy in mind. Key security aspects include:

1. **Signature-based authentication**  
   Users prove ownership of a TON wallet by signing a one-time challenge (nonce). The backend verifies the signature against the provided public key. No passwords are used, reducing the risk of credential theft.

2. **JWT tokens**  
   After successful verification, the service issues a JWT signed with its private key. These tokens have a limited lifetime and cannot be modified without invalidating the signature.

3. **Replay attack protection**  
   Each challenge is unique and valid only once. Attempting to reuse a challenge will fail, preventing replay attacks.

4. **Optional wallet existence verification**  
   The backend can optionally check that the TON wallet exists on the blockchain before issuing a token, preventing fake or invalid wallets from gaining access.

5. **No sensitive data storage**  
   The service does not store private keys or wallet secrets. Only public keys and signed challenges are transmitted temporarily for verification.

6. **Anonymity and privacy**  
   Users remain anonymous, as the service only requires wallet ownership verification. Personal data like names or emails are never collected, ensuring privacy.

These mechanisms together ensure secure, passwordless authentication while keeping users' identities protected.


## 📝 TODO List for TON OAuth Service 

Planned Features & Improvements for TON OAuth Service

- [ ] **Replay Attack Protection**
- [ ] **Retry / Rate-limiting**
- [ ] **Key Security**


## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.