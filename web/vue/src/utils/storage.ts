const storage =
  process.env.NODE_ENV == "development" ? localStorage : sessionStorage;

const tokenKey = "ddgrcf-token";
const grantKey = "ddgrcf-grant";

storage.setToken = (token: string): void => {
  storage.setItem(tokenKey, token);
};

storage.getToken = (): string | null => {
  return storage.getItem(tokenKey);
};

storage.setGrant = (token: string): void => {
  storage.setItem(grantKey, token);
};

storage.getGrant = (): string | null => {
  return storage.getItem(grantKey);
};

export default storage;
