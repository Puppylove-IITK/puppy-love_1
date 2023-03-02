import crypto from 'crypto-js';

export class Crypto {
  static hash(data) {
    const hash = crypto.SHA256(data);
    return hash.toString(crypto.enc.Hex);
  }

  static toJson(data) {
    return JSON.parse(data);
  }

  static fromJson(data) {
    return JSON.stringify(data);
  }

  static getRand(cnt = 5) {
    const randomWords = crypto.lib.WordArray.random(cnt * 4);
    return randomWords.toString(crypto.enc.Hex);
  }

  constructor(
    private password,
    private pubK,
    private priK
  ) {}

  // Key creation and storage
  // ------------------------
  newKey() {
    const key = crypto.lib.WordArray.random(32);
    this.priK = crypto.enc.Hex.parse(key.toString());
    this.pubK = this.priK;
  }

  diffieHellman(pub) {
    let pp = crypto.enc.Base64.parse(pub);
    let sharedSecret = crypto.diffieHellman(this.priK, pp);
    return sharedSecret.toString(crypto.enc.Hex);
  }

  serializePub() {
    return this.pubK.toString(crypto.enc.Base64);
  }

  deserializePub(serializedVal) {
    this.pubK = crypto.enc.Base64.parse(serializedVal);
  }

  serializePriv() {
    return this.priK.toString(crypto.enc.Base64);
  }

  deserializePriv(sec) {
    this.priK = crypto.enc.Base64.parse(sec);
  }

  // Symmetric enc and dec
  // ---------------------
  encryptSym(data) {
    const ciphertext = crypto.AES.encrypt(data, this.password);
    return ciphertext.toString();
  }

  decryptSym(data) {
    try {
      const bytes = crypto.AES.decrypt(data, this.password);
      return bytes.toString(crypto.enc.Utf8) || '{}';
    } catch (e) {
      return '{}';
    }
  }

  // Asymmetric enc and dec
  // ----------------------
  encryptAsym(data) {
    if (!this.pubK) {
      console.error('Using non-existing public key');
      return;
    }

    const ciphertext = crypto.AES.encrypt(data, this.pubK);
    return ciphertext.toString();
  }

  decryptAsym(data) {
    if (!this.priK) {
      console.error('Decryption requires private key');
      return undefined;
    }

    try {
      const bytes = crypto.AES.decrypt(data, this.priK);
      return bytes.toString(crypto.enc.Utf8);
    } catch (e) {
      return undefined;
    }
  }
}
