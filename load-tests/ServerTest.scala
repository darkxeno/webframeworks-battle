
import scala.concurrent.duration._

import io.gatling.core.Predef._
import io.gatling.http.Predef._
import io.gatling.jdbc.Predef._

class ServerTest extends Simulation {

	val httpProtocol = http
		.baseURL("http://localhost")

	val headers_0 = Map(
		"Cache-Control" -> "no-cache",
		"Origin" -> "file://"
		)

    val uri1 = "http://localhost:8181/test?limit=10"

	val scn = scenario("ServerTest")
		.exec(http("Test Request")
			.get(uri1)
			.headers(headers_0))

	setUp(scn.inject(constantUsersPerSec(300) during(1 minutes))).protocols(httpProtocol)
}