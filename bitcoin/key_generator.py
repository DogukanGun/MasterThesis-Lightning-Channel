from bitcoin import random_key, privtopub, pubtoaddr

private_key = random_key()
public_key = privtopub(private_key)
address = pubtoaddr(public_key)

print(f"Private Key: {private_key}")
print(f"Address: {address}")
