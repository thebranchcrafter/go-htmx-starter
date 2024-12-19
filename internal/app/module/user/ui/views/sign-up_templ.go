// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func SignUpContent() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body class=\"bg-gray-900 text-gray-800 font-sans scroll-smooth\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = LandingHeader().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative flex items-center justify-center min-h-screen\"><!-- Background Gradients --><div class=\"absolute inset-0 bg-gradient-to-br from-blue-600 via-purple-700 to-gray-900 opacity-80\"></div><div class=\"absolute rellax z-0 w-96 h-96 bg-gradient-to-r from-green-400 to-blue-600 rounded-full opacity-20 top-10 left-10\" data-rellax-speed=\"-6\"></div><div class=\"absolute rellax z-0 w-64 h-64 bg-gradient-to-tl from-pink-500 to-yellow-500 rounded-full opacity-30 bottom-20 right-20\" data-rellax-speed=\"3\"></div><div class=\"absolute rellax z-0 w-48 h-48 bg-gradient-to-br from-purple-400 to-indigo-600 rounded-full opacity-40 top-20 right-1/4\" data-rellax-speed=\"5\"></div><!-- Form Container --><div class=\"relative z-10 w-full max-w-md bg-white p-8 rounded-lg shadow-md\"><h2 class=\"text-2xl font-semibold text-center text-gray-800\">Create your account</h2><p class=\"text-center text-gray-500 mt-2\">Or <a href=\"/sign-in\" class=\"text-blue-500 hover:underline\">sign in to your account</a></p><form hx-post=\"/sign-up\" hx-trigger=\"submit\" hx-swap=\"none\" class=\"mt-8 space-y-4\"><div><label for=\"username\" class=\"block text-sm font-medium text-gray-700\">Username</label> <input type=\"text\" id=\"username\" name=\"username\" placeholder=\"Enter your username\" required class=\"mt-1 w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-gray-900\"></div><div><label for=\"email\" class=\"block text-sm font-medium text-gray-700\">Email address</label> <input type=\"email\" id=\"email\" name=\"email\" placeholder=\"Enter your email address\" required class=\"mt-1 w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-gray-900\"></div><div><label for=\"password\" class=\"block text-sm font-medium text-gray-700\">Password</label> <input type=\"password\" id=\"password\" name=\"password\" placeholder=\"Enter your password\" required class=\"mt-1 w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-gray-900\"></div><div><label for=\"repeat-password\" class=\"block text-sm font-medium text-gray-700\">Repeat Password</label> <input type=\"password\" id=\"repeat-password\" name=\"repeat-password\" placeholder=\"Repeat your password\" required class=\"mt-1 w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 text-gray-900\"></div><button type=\"submit\" class=\"w-full py-2 px-4 bg-blue-600 text-white font-semibold rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50\">Sign up</button></form><div id=\"alerts-container\"></div><div class=\"relative mt-6\"><div class=\"absolute inset-0 flex items-center\"><div class=\"w-full border-t border-gray-300\"></div></div><div class=\"relative flex justify-center text-sm\"><span class=\"bg-white px-2 text-gray-500\">Or continue with</span></div></div><div class=\"mt-6 flex justify-center space-x-4\"><a href=\"/auth/google\" class=\"w-full flex items-center justify-center px-8 py-3 border border-gray-300 rounded-md shadow-sm text-sm font-medium text-gray-700 bg-white hover:bg-gray-50\"><img class=\"h-6 w-6\" src=\"https://www.svgrepo.com/show/506498/google.svg\" alt=\"\"> <span class=\"ml-2\">Sign in with Google</span></a></div></div></div></body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func SignUp() templ.Component {
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
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = LandingLayout(LandingHead(), SignUpContent()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
