import styled from "styled-components";

const Large = styled.input`
  font-style: ${(props) => props.theme.f1.fontStyle};
  font-family: ${(props) => props.theme.f1.fontFamily};
  font-weight: ${(props) => props.theme.f1.fontWeight};
  font-size: ${(props) => props.theme.f1.fontSize};
  line-height: ${(props) => props.theme.f1.lineHeight};
  color: ${(props) => props.theme.f1.color};
  padding: 0.5rem 1rem;
  border: 0;
  background: var(--beach);
  outline: none;
  border-radius: 16px;
`;
const Small = styled.input`
  font-style: ${(props) => props.theme.f2.fontStyle};
  font-family: ${(props) => props.theme.f2.fontFamily};
  font-weight: ${(props) => props.theme.f2.fontWeight};
  font-size: ${(props) => props.theme.f2.fontSize};
  line-height: ${(props) => props.theme.f2.lineHeight};
  color: ${(props) => props.theme.f2.color};
  border: 0;
  padding: 0.5rem 1rem;
  border: 0;
  background: var(--background);
  outline: none;
  border-radius: 16px;
`;

export default function TextInput({ variant, value, onChange, placeholder }) {
  if (variant === "large") {
    return (
      <Large
        type="text"
        value={value}
        onChange={onChange}
        placeholder={placeholder}
      ></Large>
    );
  } else if (variant === "small") {
    return (
      <Small
        type="text"
        value={value}
        onChange={onChange}
        placeholder={placeholder}
      ></Small>
    );
  }
}
