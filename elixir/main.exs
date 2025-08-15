defmodule Gemtext do
  def convert, do: convert(false, false)

  def convert(in_pre, in_list) do
    case IO.read(:stdio, :line) do
      :eof ->
        if in_list, do: IO.puts("</ul>")
        if in_pre, do: IO.puts("</pre>")

      line ->
        cond do
          in_pre ->
            if String.trim(line) == "```" do
              IO.puts("</pre>")
              convert(false, in_list)
            else
              IO.write(line)
              convert(true, in_list)
            end
          in_list ->
            if not String.starts_with?(line, "* ") do
              IO.puts("</ul>")
              process_line(line, in_pre, false)
            else
              IO.puts("\t<li>" <> (String.trim_leading(line, "* ") |> String.trim()) <> "</li>")
              convert(in_pre, true)
            end
          true -> process_line(line, in_pre, in_list)
        end
    end
  end

  defp process_line(line, in_pre, in_list) do
    l = String.trim(line)
    cond do
      l == "" -> IO.puts("<br>"); convert(in_pre, in_list)

      String.starts_with?(l, "### ") -> IO.puts("<h3>" <> String.trim_leading(l, "### ") <> "</h3>"); convert(in_pre, in_list)
      String.starts_with?(l, "## ") -> IO.puts("<h2>" <> String.trim_leading(l, "## ") <> "</h2>"); convert(in_pre, in_list)
      String.starts_with?(l, "# ") -> IO.puts("<h1>" <> String.trim_leading(l, "# ") <> "</h1>"); convert(in_pre, in_list)

      String.starts_with?(l, "=>") ->
        rest = String.trim_leading(l, "=>") |> String.trim()
        case String.split(rest, " ", parts: 2) do
          [url] ->
            IO.puts(~s(<a href="#{url}"></a>))
          [url, text] ->
            IO.puts(~s(<a href="#{url}">#{text}</a>))
        end
        convert(in_pre, in_list)

      String.starts_with?(l, ">") ->
        IO.puts("<blockquote>" <> String.slice(l, 1..-1//1) <> "</blockquote>")
        convert(in_pre, in_list)

      String.starts_with?(l, "* ") ->
        IO.puts("<ul>")
        IO.puts("\t<li>" <> String.slice(l, 2..-1//1) <> "</li>")
        convert(in_pre, true)

      String.starts_with?(l, "```") ->
        IO.puts("<pre>")
        convert(true, in_list)

      true ->
        IO.puts("<p>" <> String.trim_trailing(line) <> "</p>")
        convert(in_pre, in_list)
    end
  end
end

Gemtext.convert()
