import React from "react";

import Card from "react-bootstrap/Card";

const ShortenResult: React.FC = () => {
  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title>Shortened URL</Card.Title>
          <Card.Text>
            The backend is implemented in Golang using a Hexagonal Archtecture.
            This architectural pattern results in clean code that is easier to
            modify and extend.
          </Card.Text>
        </Card.Body>
      </Card>
    </div>
  );
};

export default ShortenResult;
