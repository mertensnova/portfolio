// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package view

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Home() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"description\" content=\"Always laugh at the competition.\"><style>\n@import url('https://fonts.googleapis.com/css2?family=Open+Sans:ital,wght@0,300..800;1,300..800&family=Space+Mono:ital,wght@0,400;0,700;1,400;1,700&display=swap');\n</style><title>mertens</title><style>\n       body {\n            max-width: 600px;\n            margin: 10vh auto;\n            font-size: 15.5px\n            padding: 0 20px;\n            font-family: \"Space Mono\", monospace;\n            font-weight: 400;\n            font-style: normal;\n            background-color: rgb(1, 1, 1);\n            color: rgb(210, 210, 223);\n            user-select: none;\n        }\n        .social,\n        .name,\n        span{\n        color:rgb(245, 158, 11);\n        }\n        @media only screen and (max-width: 650px) {\n            body { width: 90vw !important; }\n        }\n        a {\n            color:rgb(245, 158, 11);\n            text-decoration: none;\n            user-select: text;\n        }\n        hr {\n  background-color: rgba(255, 255, 255, 0.1);\n  border-width: 0px;\n  height: 1px;\n}\n    </style></head><body><main><h2 class=\"name\">mertens</h2><hr><nav class=\"nav\"><a href=\"/\">home </a> <a href=\"/blogs\">blogs </a></nav><p></p><p>Self taught dev. Trying to learn ML, CS and Cyber Security</p><h3><span>#</span> working on:</h3><ul><li>a simple shell using C</li><li>my own grep using C++</li><li>3D renderer using C++ (not started yet)</li></ul><h3><span>#</span> blogs:</h3><ul><li><a href=\"./posts/compilers.html\">Why am I using NixOS as my daily driver</a></li></ul><br><nav><a class=\"social\" href=\"https://x.com/_blackfyre__\">[twitter]</a> <a class=\"social\" href=\"https://github.com/nyx6965\">[github]</a></nav></main></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
