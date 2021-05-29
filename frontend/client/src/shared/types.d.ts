declare module "types" {
  interface Redirect {
    url: string;
    code: string;
    created: string;
  }

  interface AppContextInterface {
    lastRedirectId: number | null;
    setLastRedirectId: (value: any) => void;
    redirects: Array<Redirect>;
    setRedirects: (value: any) => void;
  }
}

module.exports = {
  Redirect,
  AppContextInterface,
};
