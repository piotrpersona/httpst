package data

func GetHttpInfo() HTTPData {
	return HTTPData{
		Codes: []HTTPCode{
			{
				Code:        "100",
				Description: "The client SHOULD continue with its request. This interim response is\nused to inform the client that the initial part of the request has\nbeen received and has not yet been rejected by the server. The client\nSHOULD continue by sending the remainder of the request or, if the\nrequest has already been completed, ignore this response. The server\nMUST send a final response after the request has been completed. See\nsection ",
				Title:       "Continue",
			},
			{
				Code:        "101",
				Description: "The server understands and is willing to comply with the client's\nrequest, via the Upgrade message header field (section 14.42), for a\nchange in the application protocol being used on this connection. The\nserver will switch protocols to those defined by the response's\nUpgrade header field immediately after the empty line which\nterminates the 101 response.\n\nThe protocol SHOULD be switched only when it is advantageous to do\nso. For example, switching to a newer version of HTTP is advantageous\nover older versions, and switching to a real-time, synchronous\nprotocol might be advantageous when delivering resources that use\nsuch features.",
				Title:       "Switching Protocols",
			},
			{
				Code:        "200",
				Description: "The request has succeeded. The information returned with the response\nis dependent on the method used in the request, for example:\n\nGET    an entity corresponding to the requested resource is sent in\n       the response;\n\nHEAD   the entity-header fields corresponding to the requested\n       resource are sent in the response without any message-body;\n\nPOST   an entity describing or containing the result of the action;\n\nTRACE  an entity containing the request message as received by the\n       end server.",
				Title:       "OK",
			},
			{
				Code:        "201",
				Description: "The request has been fulfilled and resulted in a new resource being\ncreated. The newly created resource can be referenced by the URI(s)\nreturned in the entity of the response, with the most specific URI\nfor the resource given by a Location header field. The response\nSHOULD include an entity containing a list of resource\ncharacteristics and location(s) from which the user or user agent can\nchoose the one most appropriate. The entity format is specified by\nthe media type given in the Content-Type header field. The origin\nserver MUST create the resource before returning the 201 status code.\nIf the action cannot be carried out immediately, the server SHOULD\nrespond with 202 (Accepted) response instead.\n\nA 201 response MAY contain an ETag response header field indicating\nthe current value of the entity tag for the requested variant just\ncreated, see section ",
				Title:       "Created",
			},
			{
				Code:        "202",
				Description: "The request has been accepted for processing, but the processing has\nnot been completed.  The request might or might not eventually be\nacted upon, as it might be disallowed when processing actually takes\nplace. There is no facility for re-sending a status code from an\nasynchronous operation such as this.\n\nThe 202 response is intentionally non-committal. Its purpose is to\nallow a server to accept a request for some other process (perhaps a\nbatch-oriented process that is only run once per day) without\nrequiring that the user agent's connection to the server persist\nuntil the process is completed. The entity returned with this\nresponse SHOULD include an indication of the request's current status\nand either a pointer to a status monitor or some estimate of when the\nuser can expect the request to be fulfilled.",
				Title:       "Accepted",
			},
			{
				Code:        "203",
				Description: "The returned metainformation in the entity-header is not the\ndefinitive set as available from the origin server, but is gathered\nfrom a local or a third-party copy. The set presented MAY be a subset\nor superset of the original version. For example, including local\nannotation information about the resource might result in a superset\nof the metainformation known by the origin server. Use of this\nresponse code is not required and is only appropriate when the\nresponse would otherwise be 200 (OK).",
				Title:       "Non-Authoritative Information",
			},
			{
				Code:        "204",
				Description: "The server has fulfilled the request but does not need to return an\nentity-body, and might want to return updated metainformation. The\nresponse MAY include new or updated metainformation in the form of\nentity-headers, which if present SHOULD be associated with the\nrequested variant.\n\nIf the client is a user agent, it SHOULD NOT change its document view\nfrom that which caused the request to be sent. This response is\nprimarily intended to allow input for actions to take place without\ncausing a change to the user agent's active document view, although\nany new or updated metainformation SHOULD be applied to the document\ncurrently in the user agent's active view.\n\nThe 204 response MUST NOT include a message-body, and thus is always\nterminated by the first empty line after the header fields.",
				Title:       "No Content",
			},
			{
				Code:        "205",
				Description: "The server has fulfilled the request and the user agent SHOULD reset\nthe document view which caused the request to be sent. This response\nis primarily intended to allow input for actions to take place via\nuser input, followed by a clearing of the form in which the input is\ngiven so that the user can easily initiate another input action. The\nresponse MUST NOT include an entity.",
				Title:       "Reset Content",
			},
			{
				Code:        "206",
				Description: "The server has fulfilled the partial GET request for the resource.\nThe request MUST have included a Range header field (section 14.35)\nindicating the desired range, and MAY have included an If-Range\nheader field (section \nThe response MUST include the following header fields:\n\nIf the 206 response is the result of an If-Range request that used a\nstrong cache validator (see section 13.3.3), the response SHOULD NOT\ninclude other entity-headers. If the response is the result of an\nIf-Range request that used a weak validator, the response MUST NOT\ninclude other entity-headers; this prevents inconsistencies between\ncached entity-bodies and updated headers. Otherwise, the response\nMUST include all of the entity-headers that would have been returned\nwith a 200 (OK) response to the same request.\n\nA cache MUST NOT combine a 206 response with other previously cached\ncontent if the ETag or Last-Modified headers do not match exactly,\nsee \nA cache that does not support the Range and Content-Range headers\nMUST NOT cache 206 (Partial) responses.",
				Title:       "Partial Content",
			},
			{
				Code:        "300",
				Description: "The requested resource corresponds to any one of a set of\nrepresentations, each with its own specific location, and agent-\ndriven negotiation information (section 12) is being provided so that\nthe user (or user agent) can select a preferred representation and\nredirect its request to that location.\n\nUnless it was a HEAD request, the response SHOULD include an entity\ncontaining a list of resource characteristics and location(s) from\nwhich the user or user agent can choose the one most appropriate. The\nentity format is specified by the media type given in the Content-\nType header field. Depending upon the format and the capabilities of\n\nthe user agent, selection of the most appropriate choice MAY be\nperformed automatically. However, this specification does not define\nany standard for such automatic selection.\n\nIf the server has a preferred choice of representation, it SHOULD\ninclude the specific URI for that representation in the Location\nfield; user agents MAY use the Location field value for automatic\nredirection. This response is cacheable unless indicated otherwise.",
				Title:       "Multiple Choices",
			},
			{
				Code:        "301",
				Description: "The requested resource has been assigned a new permanent URI and any\nfuture references to this resource SHOULD use one of the returned\nURIs.  Clients with link editing capabilities ought to automatically\nre-link references to the Request-URI to one or more of the new\nreferences returned by the server, where possible. This response is\ncacheable unless indicated otherwise.\n\nThe new permanent URI SHOULD be given by the Location field in the\nresponse. Unless the request method was HEAD, the entity of the\nresponse SHOULD contain a short hypertext note with a hyperlink to\nthe new URI(s).\n\nIf the 301 status code is received in response to a request other\nthan GET or HEAD, the user agent MUST NOT automatically redirect the\nrequest unless it can be confirmed by the user, since this might\nchange the conditions under which the request was issued.",
				Title:       "Moved Permanently",
			},
			{
				Code:        "302",
				Description: "The requested resource resides temporarily under a different URI.\nSince the redirection might be altered on occasion, the client SHOULD\ncontinue to use the Request-URI for future requests.  This response\nis only cacheable if indicated by a Cache-Control or Expires header\nfield.\n\nThe temporary URI SHOULD be given by the Location field in the\nresponse. Unless the request method was HEAD, the entity of the\nresponse SHOULD contain a short hypertext note with a hyperlink to\nthe new URI(s).\n\nIf the 302 status code is received in response to a request other\nthan GET or HEAD, the user agent MUST NOT automatically redirect the\nrequest unless it can be confirmed by the user, since this might\nchange the conditions under which the request was issued.",
				Title:       "Found",
			},
			{
				Code:        "303",
				Description: "The response to the request can be found under a different URI and\nSHOULD be retrieved using a GET method on that resource. This method\nexists primarily to allow the output of a POST-activated script to\nredirect the user agent to a selected resource. The new URI is not a\nsubstitute reference for the originally requested resource. The 303\nresponse MUST NOT be cached, but the response to the second\n(redirected) request might be cacheable.\n\nThe different URI SHOULD be given by the Location field in the\nresponse. Unless the request method was HEAD, the entity of the\nresponse SHOULD contain a short hypertext note with a hyperlink to\nthe new URI(s).",
				Title:       "See Other",
			},
			{
				Code:        "304",
				Description: "If the client has performed a conditional GET request and access is\nallowed, but the document has not been modified, the server SHOULD\nrespond with this status code. The 304 response MUST NOT contain a\nmessage-body, and thus is always terminated by the first empty line\nafter the header fields.\n\nThe response MUST include the following header fields:\n\nIf a clockless origin server obeys these rules, and proxies and\nclients add their own Date to any response received without one (as\nalready specified by [RFC 2068], section \nIf the conditional GET used a strong cache validator (see section\n13.3.3), the response SHOULD NOT include other entity-headers.\nOtherwise (i.e., the conditional GET used a weak validator), the\nresponse MUST NOT include other entity-headers; this prevents\ninconsistencies between cached entity-bodies and updated headers.\n\nIf a 304 response indicates an entity not currently cached, then the\ncache MUST disregard the response and repeat the request without the\nconditional.\n\nIf a cache uses a received 304 response to update a cache entry, the\ncache MUST update the entry to reflect any new field values given in\nthe response.",
				Title:       "Not Modified",
			},
			{
				Code:        "305",
				Description: "The requested resource MUST be accessed through the proxy given by\nthe Location field. The Location field gives the URI of the proxy.\nThe recipient is expected to repeat this single request via the\nproxy. 305 responses MUST only be generated by origin servers.",
				Title:       "Use Proxy",
			},
			{
				Code:        "306",
				Description: "The 306 status code was used in a previous version of the\nspecification, is no longer used, and the code is reserved.",
				Title:       "(Unused)",
			},
			{
				Code:        "307",
				Description: "The requested resource resides temporarily under a different URI.\nSince the redirection MAY be altered on occasion, the client SHOULD\ncontinue to use the Request-URI for future requests.  This response\nis only cacheable if indicated by a Cache-Control or Expires header\nfield.\n\nThe temporary URI SHOULD be given by the Location field in the\nresponse. Unless the request method was HEAD, the entity of the\nresponse SHOULD contain a short hypertext note with a hyperlink to\nthe new URI(s) , since many pre-HTTP/1.1 user agents do not\nunderstand the 307 status. Therefore, the note SHOULD contain the\ninformation necessary for a user to repeat the original request on\nthe new URI.\n\nIf the 307 status code is received in response to a request other\nthan GET or HEAD, the user agent MUST NOT automatically redirect the\nrequest unless it can be confirmed by the user, since this might\nchange the conditions under which the request was issued.",
				Title:       "Temporary Redirect",
			},
			{
				Code:        "400",
				Description: "The request could not be understood by the server due to malformed\nsyntax. The client SHOULD NOT repeat the request without\nmodifications.",
				Title:       "Bad Request",
			},
			{
				Code:        "401",
				Description: "The request requires user authentication. The response MUST include a\nWWW-Authenticate header field (section 14.47) containing a challenge\napplicable to the requested resource. The client MAY repeat the\nrequest with a suitable Authorization header field (section ",
				Title:       "Unauthorized",
			},
			{
				Code:        "402",
				Description: "This code is reserved for future use.",
				Title:       "Payment Required",
			},
			{
				Code:        "403",
				Description: "The server understood the request, but is refusing to fulfill it.\nAuthorization will not help and the request SHOULD NOT be repeated.\nIf the request method was not HEAD and the server wishes to make\npublic why the request has not been fulfilled, it SHOULD describe the\nreason for the refusal in the entity.  If the server does not wish to\nmake this information available to the client, the status code 404\n(Not Found) can be used instead.",
				Title:       "Forbidden",
			},
			{
				Code:        "404",
				Description: "The server has not found anything matching the Request-URI. No\nindication is given of whether the condition is temporary or\npermanent. The 410 (Gone) status code SHOULD be used if the server\nknows, through some internally configurable mechanism, that an old\nresource is permanently unavailable and has no forwarding address.\nThis status code is commonly used when the server does not wish to\nreveal exactly why the request has been refused, or when no other\nresponse is applicable.",
				Title:       "Not Found",
			},
			{
				Code:        "405",
				Description: "The method specified in the Request-Line is not allowed for the\nresource identified by the Request-URI. The response MUST include an\nAllow header containing a list of valid methods for the requested\nresource.",
				Title:       "Method Not Allowed",
			},
			{
				Code:        "406",
				Description: "The resource identified by the request is only capable of generating\nresponse entities which have content characteristics not acceptable\naccording to the accept headers sent in the request.\n\nUnless it was a HEAD request, the response SHOULD include an entity\ncontaining a list of available entity characteristics and location(s)\nfrom which the user or user agent can choose the one most\nappropriate. The entity format is specified by the media type given\nin the Content-Type header field. Depending upon the format and the\ncapabilities of the user agent, selection of the most appropriate\nchoice MAY be performed automatically. However, this specification\ndoes not define any standard for such automatic selection.\n\nIf the response could be unacceptable, a user agent SHOULD\ntemporarily stop receipt of more data and query the user for a\ndecision on further actions.",
				Title:       "Not Acceptable",
			},
			{
				Code:        "407",
				Description: "This code is similar to 401 (Unauthorized), but indicates that the\nclient must first authenticate itself with the proxy. The proxy MUST\nreturn a Proxy-Authenticate header field (section ",
				Title:       "Proxy Authentication Required",
			},
			{
				Code:        "408",
				Description: "The client did not produce a request within the time that the server\nwas prepared to wait. The client MAY repeat the request without\nmodifications at any later time.",
				Title:       "Request Timeout",
			},
			{
				Code:        "409",
				Description: "The request could not be completed due to a conflict with the current\nstate of the resource. This code is only allowed in situations where\nit is expected that the user might be able to resolve the conflict\nand resubmit the request. The response body SHOULD include enough\n\ninformation for the user to recognize the source of the conflict.\nIdeally, the response entity would include enough information for the\nuser or user agent to fix the problem; however, that might not be\npossible and is not required.\n\nConflicts are most likely to occur in response to a PUT request. For\nexample, if versioning were being used and the entity being PUT\nincluded changes to a resource which conflict with those made by an\nearlier (third-party) request, the server might use the 409 response\nto indicate that it can't complete the request. In this case, the\nresponse entity would likely contain a list of the differences\nbetween the two versions in a format defined by the response\nContent-Type.",
				Title:       "Conflict",
			},
			{
				Code:        "410",
				Description: "The requested resource is no longer available at the server and no\nforwarding address is known. This condition is expected to be\nconsidered permanent. Clients with link editing capabilities SHOULD\ndelete references to the Request-URI after user approval. If the\nserver does not know, or has no facility to determine, whether or not\nthe condition is permanent, the status code 404 (Not Found) SHOULD be\nused instead. This response is cacheable unless indicated otherwise.\n\nThe 410 response is primarily intended to assist the task of web\nmaintenance by notifying the recipient that the resource is\nintentionally unavailable and that the server owners desire that\nremote links to that resource be removed. Such an event is common for\nlimited-time, promotional services and for resources belonging to\nindividuals no longer working at the server's site. It is not\nnecessary to mark all permanently unavailable resources as \"gone\" or\nto keep the mark for any length of time -- that is left to the\ndiscretion of the server owner.",
				Title:       "Gone",
			},
			{
				Code:        "411",
				Description: "The server refuses to accept the request without a defined Content-\nLength. The client MAY repeat the request if it adds a valid\nContent-Length header field containing the length of the message-body\nin the request message.",
				Title:       "Length Required",
			},
			{
				Code:        "412",
				Description: "The precondition given in one or more of the request-header fields\nevaluated to false when it was tested on the server. This response\ncode allows the client to place preconditions on the current resource\nmetainformation (header field data) and thus prevent the requested\nmethod from being applied to a resource other than the one intended.",
				Title:       "Precondition Failed",
			},
			{
				Code:        "413",
				Description: "The server is refusing to process a request because the request\nentity is larger than the server is willing or able to process. The\nserver MAY close the connection to prevent the client from continuing\nthe request.\n\nIf the condition is temporary, the server SHOULD include a Retry-\nAfter header field to indicate that it is temporary and after what\ntime the client MAY try again.",
				Title:       "Request Entity Too Large",
			},
			{
				Code:        "414",
				Description: "The server is refusing to service the request because the Request-URI\nis longer than the server is willing to interpret. This rare\ncondition is only likely to occur when a client has improperly\nconverted a POST request to a GET request with long query\ninformation, when the client has descended into a URI \"black hole\" of\nredirection (e.g., a redirected URI prefix that points to a suffix of\nitself), or when the server is under attack by a client attempting to\nexploit security holes present in some servers using fixed-length\nbuffers for reading or manipulating the Request-URI.",
				Title:       "Request-URI Too Long",
			},
			{
				Code:        "415",
				Description: "The server is refusing to service the request because the entity of\nthe request is in a format not supported by the requested resource\nfor the requested method.",
				Title:       "Unsupported Media Type",
			},
			{
				Code:        "416",
				Description: "A server SHOULD return a response with this status code if a request\nincluded a Range request-header field (section 14.35), and none of\nthe range-specifier values in this field overlap the current extent\nof the selected resource, and the request did not include an If-Range\nrequest-header field. (For byte-ranges, this means that the first-\nbyte-pos of all of the byte-range-spec values were greater than the\ncurrent length of the selected resource.)\n\nWhen this status code is returned for a byte-range request, the\nresponse SHOULD include a Content-Range entity-header field\nspecifying the current length of the selected resource (see section",
				Title:       "Requested Range Not Satisfiable",
			},
			{
				Code:        "417",
				Description: "The expectation given in an Expect request-header field (see section\n14.20) could not be met by this server, or, if the server is a proxy,\nthe server has unambiguous evidence that the request could not be met\nby the next-hop server.",
				Title:       "Expectation Failed",
			},
			{
				Code:        "500",
				Description: "The server encountered an unexpected condition which prevented it\nfrom fulfilling the request.",
				Title:       "Internal Server Error",
			},
			{
				Code:        "501",
				Description: "The server does not support the functionality required to fulfill the\nrequest. This is the appropriate response when the server does not\nrecognize the request method and is not capable of supporting it for\nany resource.",
				Title:       "Not Implemented",
			},
			{
				Code:        "502",
				Description: "The server, while acting as a gateway or proxy, received an invalid\nresponse from the upstream server it accessed in attempting to\nfulfill the request.",
				Title:       "Bad Gateway",
			},
			{
				Code:        "503",
				Description: "The server is currently unable to handle the request due to a\ntemporary overloading or maintenance of the server. The implication\nis that this is a temporary condition which will be alleviated after\nsome delay. If known, the length of the delay MAY be indicated in a\nRetry-After header. If no Retry-After is given, the client SHOULD\nhandle the response as it would for a 500 response.",
				Title:       "Service Unavailable",
			},
			{
				Code:        "504",
				Description: "The server, while acting as a gateway or proxy, did not receive a\ntimely response from the upstream server specified by the URI (e.g.\nHTTP, FTP, LDAP) or some other auxiliary server (e.g. DNS) it needed\nto access in attempting to complete the request.",
				Title:       "Gateway Timeout",
			},
		},
		Groups: []HTTPGroup{
			{
				Description: "This class of status code indicates a provisional response,\nconsisting only of the Status-Line and optional headers, and is\nterminated by an empty line. There are no required headers for this\nclass of status code. Since HTTP/1.0 did not define any 1xx status\ncodes, servers MUST NOT send a 1xx response to an HTTP/1.0 client\nexcept under experimental conditions.\n\nA client MUST be prepared to accept one or more 1xx status responses\nprior to a regular response, even if the client does not expect a 100\n(Continue) status message. Unexpected 1xx status responses MAY be\nignored by a user agent.\n\nProxies MUST forward 1xx responses, unless the connection between the\nproxy and its client has been closed, or unless the proxy itself\nrequested the generation of the 1xx response. (For example, if a\n\nproxy adds a \"Expect: 100-continue\" field when it forwards a request,\nthen it need not forward the corresponding 100 (Continue)\nresponse(s).)",
				Prefix:      "1xx",
				Title:       "Informational 1xx",
			},
			{
				Description: "This class of status code indicates that the client's request was\nsuccessfully received, understood, and accepted.",
				Prefix:      "2xx",
				Title:       "Successful 2xx",
			},
			{
				Description: "This class of status code indicates that further action needs to be\ntaken by the user agent in order to fulfill the request.  The action\nrequired MAY be carried out by the user agent without interaction\nwith the user if and only if the method used in the second request is\nGET or HEAD. A client SHOULD detect infinite redirection loops, since\nsuch loops generate network traffic for each redirection.",
				Prefix:      "3xx",
				Title:       "Redirection 3xx",
			},
			{
				Description: "The 4xx class of status code is intended for cases in which the\nclient seems to have erred. Except when responding to a HEAD request,\nthe server SHOULD include an entity containing an explanation of the\nerror situation, and whether it is a temporary or permanent\ncondition. These status codes are applicable to any request method.\nUser agents SHOULD display any included entity to the user.\n\nIf the client is sending data, a server implementation using TCP\nSHOULD be careful to ensure that the client acknowledges receipt of\nthe packet(s) containing the response, before the server closes the\ninput connection. If the client continues sending data to the server\nafter the close, the server's TCP stack will send a reset packet to\nthe client, which may erase the client's unacknowledged input buffers\nbefore they can be read and interpreted by the HTTP application.",
				Prefix:      "4xx",
				Title:       "Client Error 4xx",
			},
			{
				Description: "Response status codes beginning with the digit 5 indicate cases in which the server is aware that it has erred or is incapable of performing the request.\nExcept when responding to a HEAD request, the server SHOULD include an entity containing an explanation of the error situation, and whether it is a temporary or permanent condition.\nUser agents SHOULD display any included entity to the user. These response codes are applicable to any request method.",
				Prefix:      "5xx",
				Title:       "Server Error 5xx",
			},
		},
	}
}
