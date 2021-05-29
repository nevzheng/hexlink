declare module "types" {
  interface Redirect {
    url: string;
    code: string;
    created: string;
  }

  interface ContextProps {
    lastRedirectId: number;
    redirects: Array<Redirect>;
  }
}

module.exports = {
  Redirect,
  ContextProps,
};
