import styled from "styled-components";
import { Link } from "react-router-dom";
const StyledLoginButton = styled.div`
  font-style: ${(props) => props.theme.h3.fontStyle};
  font-family: ${(props) => props.theme.h3.fontFamily};
  font-weight: ${(props) => props.theme.h3.fontWeight};
  font-size: ${(props) => props.theme.h3.fontSize};
  line-height: ${(props) => props.theme.h3.lineHeight};
  color: ${(props) => props.theme.h3.color};
  padding: 0 2rem;
`;
export default function LoginButton({ children, href }) {
  return (
    <Link to={href} style={{textDecoration: "none"}}>
      <StyledLoginButton>{children}</StyledLoginButton>
    </Link>
  );
}
