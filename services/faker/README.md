# Faker

Generates fake data for testing purposes.

`GET /feed/{uuidv4}` - generates a deterministic "fake" RSS feed with a random number of articles. subsequent requests will return the same channel, but with different items.

## Tilt integration

There is a [button](https://github.com/ericbutera/amalgam/blob/ad3d79839030889826a8fb2f0c0dcad48bf9d06e/Tiltfile#L83-L88) in the Tilt UI that will generate a fake feed.

## Example

```xml
<rss version="2.0">
  <channel>
    <title>secured line</title>
    <link>https://faker:8080/feed/81ade5f1-e5fe-499e-b197-1928f8f68a6f</link>
    <description>Everybody regularly is.</description>
    <item>
      <title>Trip early without ours phew.</title>
      <link>http://www.nationalmaximize.biz/disintermediate/revolutionize</link>
      <description>Talented lie quarterly why her. Of always hardly also niche. Additionally ouch up is hundreds. His whose African down whose. Relaxation including fact conclude on. Does less did melon archipelago.</description>
      <pubDate>Sat, 16 Nov 2024 01:57:11 UTC</pubDate>
    </item>
    <item>
      <title>Her result remind anybody should.</title>
      <link>http://www.leadtechnologies.com/transparent/virtual</link>
      <description>Bunch film which troupe all. Possess over perfectly you forgive. E.g. had those modern hourly. Heap has dream indeed as. Time caused basket cackle late. Often hence nobody everybody these.</description>
      <pubDate>Sat, 16 Nov 2024 00:57:11 UTC</pubDate>
    </item>
  </channel>
</rss>
```
