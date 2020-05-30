var CACHE_NAME = 'mathch-cache-v1';
var urlsToCache = [
    '/',
    '/css/m.css',
    '/js/jquery-3.5.1.min.js',
    '/js/m.js',
    '/img/match0.webp',
    '/img/match1.webp',
    '/img/match2.webp',
    '/img/match3.webp',
    '/img/match4.webp',
    '/img/match5.webp',
    '/img/match6.webp',
    '/img/match7.webp',
    '/img/match8.webp',
    '/img/match9.webp',
    '/img/matchplus.webp',
    '/img/matchsub.webp',
    '/img/matchdiv.webp',
    '/img/matchmul.webp',
    '/img/matcheq.webp',
    '/img/match0.jp2',
    '/img/match1.jp2',
    '/img/match2.jp2',
    '/img/match3.jp2',
    '/img/match4.jp2',
    '/img/match5.jp2',
    '/img/match6.jp2',
    '/img/match7.jp2',
    '/img/match8.jp2',
    '/img/match9.jp2',
    '/img/matchplus.jp2',
    '/img/matchsub.jp2',
    '/img/matchdiv.jp2',
    '/img/matchmul.jp2',
    '/img/matcheq.jp2',
    '/img/icon-192.png',
    '/img/icon-512.png',
    '/manifest.webmanifest',
    '/favicon.ico',
    '/favicon-16x16.png',
    '/favicon-32x32.png',
];

self.addEventListener('install', function(event) {
    // Perform install steps
    event.waitUntil(
        caches.open(CACHE_NAME)
            .then(function(cache) {
                console.log('Opened cache');
                return cache.addAll(urlsToCache);
            })
    );
});

self.addEventListener('fetch', function(event) {
    event.respondWith(
        caches.match(event.request)
            .then(function(response) {
                    // Cache hit - return response
                    if (response) {
                        return response;
                    }
                    return fetch(event.request);
                }
            )
    );
});