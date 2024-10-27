from fastapi import FastAPI, Response, HTTPException, Request
from fastapi.middleware.cors import CORSMiddleware
from fastapi.responses import RedirectResponse, JSONResponse
from pydantic import BaseModel
from typing import Optional
import requests
import json


app = FastAPI(
    title="SSO API for PQC Backend",
    description=(
        "**Qdrant Database**\n"
        "   - **URL**: [http://qsign.southeastasia.cloudapp.azure.com:6333/dashboard](http://qsign.southeastasia.cloudapp.azure.com:6333/dashboard)\n"
    ),
    version="1.0.0",
)

# Custom Exception Handler
@app.exception_handler(HTTPException)
async def custom_http_exception_handler(request: Request, exc: HTTPException):
    return JSONResponse(
        status_code=exc.status_code,
        content={"error": exc.detail}
    )

# Configure CORS to allow all origins
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Allows all origins
    allow_credentials=True,
    allow_methods=["*"],  # Allows all HTTP methods (GET, POST, etc.)
    allow_headers=["*"],  # Allows all headers
)

backend_url = "http://localhost:4444"  # Replace with actual backend URL if necessary

# Pydantic models for request and response bodies


class RegisterRequest(BaseModel):
    kyberPublicKey: str
    username: str
    email: str


class RegisterResponse(BaseModel):
    message: str


class LoginRequest(BaseModel):
    kyberPublicKey: str
    clientId: Optional[str] = None  # clientId is optional


class LoginResponse(BaseModel):
    accessToken: str
    redirectUrl: str


class ValidateRequest(BaseModel):
    jwtToken: str
    clientId: Optional[str] = None


class ValidateResponse(BaseModel):
    isValid: bool
    isExpired: bool


class KeyPairResponse(BaseModel):
    publicKey: str
    privateKey: str


class ClientRegistrationRequest(BaseModel):
    clientName: str
    redirectUrl: str


class ClientRegistrationResponse(BaseModel):
    clientId: str
    publicKey: str
    privateKey: str


# Helper function to relay requests to the backend
def relay_request(endpoint: str, data: dict = None, method: str = "POST"):
    headers = {"Content-Type": "application/json"}

    if method == "POST":
        backend_response = requests.post(
            url=f"{backend_url}{endpoint}", headers=headers, data=json.dumps(data)
        )
    else:  # For GET requests
        backend_response = requests.get(url=f"{backend_url}{endpoint}", headers=headers)

    return Response(
        content=backend_response.content,
        status_code=backend_response.status_code,
        headers=dict(backend_response.headers),
    )


@app.post(
    "/register",
    tags=["User Registration"],
    summary="Register a new user",
    response_model=RegisterResponse,
)
async def register(request_data: RegisterRequest):
    """
    Registers a new user with a public key, username, and email.
    Relays the request to the backend.
    """
    return relay_request("/register", data=request_data.dict())


@app.post(
    "/login", tags=["User Login"], summary="User login", response_model=LoginResponse
)
async def login(request: Request, request_data: LoginRequest):
    """
    Logs in a user using their public key and optionally a client ID.
    Relays the request to the backend and returns a JWT token and redirect URL.
    """

    client_ip = request.headers.get('cf-connecting-ip') or request.client.host
    login_restriction(client_ip)

    return relay_request("/login", data=request_data.dict())


@app.post(
    "/validate",
    tags=["JWT Validation"],
    summary="Validate JWT token",
    response_model=ValidateResponse,
)
async def validate(request_data: ValidateRequest):
    """
    Validates a JWT token using the user's public key or optionally a client ID.
    Relays the request to the backend and returns whether the token is valid and if it has expired.
    """
    return relay_request("/validate", data=request_data.dict())


@app.get(
    "/generate/dilithium",
    tags=["Key Generation"],
    summary="Generate Dilithium key pair",
    response_model=KeyPairResponse,
)
async def generate_dilithium_key_pair():
    """
    Generates a Dilithium key pair and returns the public and private keys.
    Relays the request to the backend.
    """
    return relay_request("/generate/dilithium", method="GET")


@app.get(
    "/generate/kyber",
    tags=["Key Generation"],
    summary="Generate Kyber key pair",
    response_model=KeyPairResponse,
)
async def generate_kyber_key_pair():
    """
    Generates a Kyber key pair and returns the public and private keys.
    Relays the request to the backend.
    """
    return relay_request("/generate/kyber", method="GET")


@app.post(
    "/client-registration",
    tags=["Client Registration"],
    summary="Register a new client",
    response_model=ClientRegistrationResponse,
)
async def client_registration(request_data: ClientRegistrationRequest):
    """
    Registers a new client and generates a Dilithium key pair for the client.
    Relays the request to the backend and returns the client ID, public key, and private key.
    """
    return relay_request("/client-registration", data=request_data.dict())


@app.get("/health", tags=["Health Check"], summary="API Health Check")
async def health_check():
    """
    Basic health check for the API.
    """
    return {"status": "Running"}


# Redirect user to /docs if they visit /
@app.get("/", include_in_schema=False)
async def redirect_to_docs():
    return RedirectResponse(url="/docs")


def login_restriction(ip: str, allowed_country: str = "LK"):

    # Allow if IP is local
    if ip in ["localhost", "127.0.0.1"]:
        return True

    response = requests.get(f"https://ipapi.co/{ip}/country/")

    if response.status_code != 200:
        raise HTTPException(
            status_code=500, detail="Unable to determine country from IP"
        )

    user_country = response.text.strip()
    if user_country != allowed_country:
        raise HTTPException(
            status_code=403,
            detail=f"Risky login attempt detected",
        )


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(app, host="0.0.0.0", port=8000)
