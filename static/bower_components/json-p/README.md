json-p
======
This element makes JSONP requests.  It exposes a similar API as [core-ajax](https://www.polymer-project.org/docs/elements/core-elements.html#core-ajax).

See the [component page](https://github.com/404) for more information.

## Installation

Install the package
```
bower install json-p --save
```

## Usage

Use the following markup on your page
```
<json-p auto
        url="https://archive.org/advancedsearch.php?q=collection%3A(GratefulDead%20AND%20etree)&rows=5&page=0&output=json"
        on-jsonp-complete="{{handleResponse}}"
        response="{{response}}
</json-p>
```

## Attributes

*auto* - boolean - If true, automatically performs an Ajax request on-ready or when url changes. Defaults to false.

*url* - string - The URL target of the request. Defaults to ''.

## Methods

*go* - Method which performs the server request.
