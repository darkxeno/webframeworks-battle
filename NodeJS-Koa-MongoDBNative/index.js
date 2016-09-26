
var koa = require('koa');
var route = require('koa-route');
var mongodb = require('mongodb');
var Promise = require('bluebird');

var app = koa();
var collection;

// Connection url
const mongoURL = process.env.MONGODB_URL || 'localhost:27017';
const url = 'mongodb://'+mongoURL+'/local';
// Connect using MongoClient
mongodb.MongoClient.connect(url, { promiseLibrary: Promise })
  .then((db)=>{
    collection = db.collection('testData');   
  });

app.use(route.get('/',function *home(next){
  this.body = 'NodeJS-Koa-MongoDBNative';
}));

app.use(route.get('/test',function *test(next){
  if(collection){
    const limit = parseInt(this.query.limit || 100);
    const results = yield collection.find().limit(limit).toArray();
    this.body={data: results};    
  } else {
    this.throw(404,'No collection',{ error:'No collection' });
  }
}));

app.on('error', function(err, ctx){
  console.error('server error', err);
});

app.listen(8181);
