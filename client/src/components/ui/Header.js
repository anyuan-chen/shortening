import styled from "styled-components";
const H1 = styled.h1`
    font-style: ${(props) => props.theme.h1.fontStyle};
    font-family: ${(props) => props.theme.h1.fontFamily};
    font-weight: ${(props) => props.theme.h1.fontWeight};
    font-size: ${(props) => props.theme.h1.fontSize};
    line-height: ${(props) => props.theme.h1.lineHeight};
    color: ${(props) => props.theme.h1.color};
`;
const H2 = styled.h2`
    font-style: ${(props) => props.theme.h2.fontStyle};
    font-family: ${(props) => props.theme.h2.fontFamily};
    font-weight: ${(props) => props.theme.h2.fontWeight};
    font-size: ${(props) => props.theme.h2.fontSize};
    line-height: ${(props) => props.theme.h2.lineHeight};
    color: ${(props) => props.theme.h2.color};
`;

export default function Header({ children, type }) {
  if (type === "h1") {
    return <H1>{children}</H1>;
  } else if (type === "h2") {
    return <H2>{children}</H2>;
  }
  return <h1>not configured properly!</h1>;
}
