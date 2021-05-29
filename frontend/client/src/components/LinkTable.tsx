import React from "react";
import Table from "react-bootstrap/Table";

import { Redirect } from "types";

import LinkTableRow from "./LinkTableRow";

interface LinkTableProps {
  redirects: Array<Redirect>;
}

const LinkTable: React.FC<LinkTableProps> = ({ redirects }: LinkTableProps) => {
  return (
    <div>
      <h3>Results</h3>
      <p>
        Links Not Persisted on Refresh. Would need to implement user or
        anonymous sessions
      </p>
      <Table size="sm">
        <thead>
          <th>#</th>
          <th>Full URL</th>
          <th>Shortened</th>
        </thead>
        <tbody>
          {redirects.map((redirect, index) => {
            return <LinkTableRow redirect={redirect} key={index} />;
          })}
        </tbody>
      </Table>
    </div>
  );
};

export default LinkTable;
