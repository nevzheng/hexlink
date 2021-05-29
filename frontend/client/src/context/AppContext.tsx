import React, { useState } from "react";
import { AppContextInterface } from "types";

export const AppContext = React.createContext<AppContextInterface | null>(null);

const AppContextProvider: React.FC = (props) => {
  const [lastRedirectId, setLastRedirectId] = useState(null);
  const [redirects, setRedirects] = useState([]);
  return (
    <AppContext.Provider
      value={{
        lastRedirectId,
        setLastRedirectId,
        redirects,
        setRedirects,
      }}
    >
      {
        // TODO: Remove workaround
        // eslint-disable-next-line react/prop-types
        props.children
      }
    </AppContext.Provider>
  );
};

export default AppContextProvider;
