
var koa = require('koa');
var route = require('koa-route');
var mongodb = require('mongodb');
var Promise = require('bluebird');

var app = koa();
var collection;

const mongoConfig = {
  promiseLibrary: Promise,
  replSet: {
    w: 1,
    wtimeout: 0,
    fsync: false,
    journal: false,
    readPreference: null,
    nativeParser: false,
    forceServerObjectId: false,
    recordQueryStats: false,
    retryMiliSeconds: 5000,
    numberOfRetries: 5,  
    
    poolSize: 20,
      socketOptions: {
        noDelay: true,
        keepAlive: 0,
        connectTimeoutMS: 0,
        socketTimeoutMS: 0
      }
  }
};

// Connection url
const url = 'mongodb://localhost:27017/local';
// Connect using MongoClient
mongodb.MongoClient.connect(url, mongoConfig)
  .then((db)=>{
    collection = db.collection('testData');   
  });

app.use(route.get('/',function *home(next){
  this.body = 'hello world';
}));

app.use(route.get('/test',function *test(next){
  if(collection){
    const limit = this.query.limit || 100;
    const results = yield collection.find().limit(10).toArray();
    this.body={data: results};    
  } else {
    this.throw(404,'No collection',{ error:'No collection' });
  }
}));

app.on('error', function(err, ctx){
  console.error('server error', err);
});

app.listen(8181);
