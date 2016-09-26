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
import MongoDB

func mongoHandler(request: HTTPRequest, _ response: HTTPResponse){
    

    let mongoURL = getEnvVar(name:"MONGODB_URL")=="" ? "mongodb://localhost" : getEnvVar(name:"MONGODB_URL");

    // open a connection
    let client = try! MongoClient(uri: mongoURL)
    
    // set database, assuming "test" exists
    let db = client.getDatabase(name: "local")
    
    // define collection
    guard let collection = db.getCollection(name: "testData") else {
        print("testData not found")
        return
    }   
    
    // Here we clean up our connection,
    // by backing out in reverse order created
    defer {
        collection.close()
        db.close()
        client.close()
    }

    // Perform a "find" on the perviously defined collection
    let fnd = collection.find(query: BSON(), limit: Int( request.param(name: "limit") ?? "0" ) ?? 0 )
    
    // Initialize empty array to receive formatted results
    var arr = [String]()
    
    // The "fnd" cursor is typed as MongoCursor, which is iterable
    for x in fnd! {
        arr.append(x.asString)
    }

    // return a formatted JSON array.
    let returning = "{\"data\":[\(arr.joined(separator: ","))]}"

    // Return the JSON string
    response.setHeader(.contentType, value: "application/json; charset=utf-8")
    response.appendBody(string: returning)
    response.completed()
}
