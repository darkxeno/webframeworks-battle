//
//  mongoHandler.swift
//  PerfectTemplate
//
//  Created by Constantino Fernandez Traba on 24/9/16.
//
//

import PerfectLib
import PerfectHTTP
import PerfectHTTPServer

func mongoHandler(request: HTTPRequest, _ response: HTTPResponse){
    
    let returning: String;

    do {
        let fnd = try testCollection.find(limitedTo: Int32( request.param(name: "limit") ?? "0" ) ?? 0)
        let arr = Array(fnd);
    
        // return a formatted JSON array.
        returning = "{\"data\":[\(arr)]}"

    } catch {
        returning = "{\"error\": \"Error executing mongodb find\"}"
    }
    
    // Return the JSON string
    response.appendBody(string: returning)
    response.completed()
}
