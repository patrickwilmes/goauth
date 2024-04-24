#include <stdio.h>
#include <openssl/rand.h>

#define KEY_LENGTH 32  

int main() {
    unsigned char key[KEY_LENGTH];

    if (RAND_bytes(key, sizeof(key)) != 1) {
        fprintf(stderr, "Error: Failed to generate random bytes\n");
        return 1;
    }

    printf("Generated Secret Key (hexadecimal):\n");
    for (int i = 0; i < KEY_LENGTH; i++) {
        printf("%02x", key[i]);
    }
    printf("\n");

    OPENSSL_cleanse(key, sizeof(key));

    return 0;
}

