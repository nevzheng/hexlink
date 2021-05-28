import React from "react";
import Table from "react-bootstrap/Table";

import LinkTableRow from "./LinkTableRow";

export const LinkTable = () => {
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
          <LinkTableRow />
          <LinkTableRow />
          <LinkTableRow />
        </tbody>
      </Table>
    </div>
  );
};

export default LinkTable;
