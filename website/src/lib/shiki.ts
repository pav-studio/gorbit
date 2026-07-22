import { createHighlighter } from "shiki";

const highlighterPromise = createHighlighter({
    themes: [
        "github-dark",
        "github-light"
    ],
    langs: [
        "go",
        "javascript",
        "typescript",
        "json",
        "bash"
    ]
});

export async function highlight(
    code: string,
    lang: string = "text",
    dark = "light"
) {
    const highlighter = await highlighterPromise;

    return highlighter.codeToHtml(code, {
        lang,
        theme: dark==="light" ? "github-light" : "github-dark",
    });
}