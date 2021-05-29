import { Redirect } from "types";

interface LinkTableRowProp {
  redirect: Redirect;
  key: number;
}

const LinkTableRow: React.FC<LinkTableRowProp> = (props: LinkTableRowProp) => {
  return (
    <tr>
      <td>{props.key}</td>
      <td>{props.redirect.url}</td>
      <td>{props.redirect.code}</td>
    </tr>
  );
};

export default LinkTableRow;
