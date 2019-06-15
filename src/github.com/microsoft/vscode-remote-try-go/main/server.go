/*----------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See LICENSE in the project root for license information.
 *---------------------------------------------------------------------------------------*/

package main

import (
	"fmt"
	"github.com/microsoft/vscode-remote-try-go/hello"
	"io"
	"net/http"
)

func helloOld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, hello.Hello())
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", helloOld)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}