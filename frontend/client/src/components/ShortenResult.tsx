import Card from "react-bootstrap/Card";

interface ShortenedResultProps {
  url?: string;
  shortUrl?: string;
}

const ShortenResult: React.FC<ShortenedResultProps> = ({
  url,
  shortUrl,
}: ShortenedResultProps) => {
  return (
    <div>
      <Card>
        <Card.Body>
          <Card.Title>Result</Card.Title>
          <Card.Text>
            {url ? url + "=>" + shortUrl : "Please Shorten a URL"}
          </Card.Text>
        </Card.Body>
      </Card>
    </div>
  );
};

export default ShortenResult;
