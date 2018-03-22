{-# LANGUAGE OverloadedStrings #-}
module FirstApp.Main (runApp) where

import           Network.Wai              (Application, Request, Response,
                                           pathInfo, requestMethod, responseLBS,
                                           strictRequestBody)
import           Network.Wai.Handler.Warp (run)

import           Network.HTTP.Types       (Status, hContentType,
                                           status200, status400, status404)

import qualified Data.ByteString.Lazy     as LBS

import           Data.Either              (either)

import           Data.Text                (Text)
import           Data.Text.Encoding       (decodeUtf8)

import           FirstApp.Types           (ContentType(..), Error(..), RqType(..),
                                           mkCommentText, mkTopic,
                                           renderContentType)

-- --------------------------------------------
-- - Don't start here, go to FirstApp.Types!  -
-- --------------------------------------------

-- | Some helper functions to make our lives a little more DRY.
mkResponse
  :: Status
  -> ContentType
  -> LBS.ByteString
  -> Response
mkResponse status content txt = responseLBS
  status
  [("Content-Type", renderContentType content)]
  txt

resp200
  :: ContentType
  -> LBS.ByteString
  -> Response
resp200 = mkResponse status200

resp404
  :: ContentType
  -> LBS.ByteString
  -> Response
resp404 = mkResponse status404

resp400
  :: ContentType
  -> LBS.ByteString
  -> Response
resp400 = mkResponse status400

-- These next few functions will take raw request information and construct one
-- of our types.
mkAddRequest
  :: Text
  -> LBS.ByteString
  -> Either Error RqType
mkAddRequest topicTxt comentTxt =
  case mkTopic topicTxt of
    Right topic -> case mkCommentText $ lazyByteStringToStrictText comentTxt of
      Right comment -> Right $ AddRq topic comment
      Left err -> Left err
    Left err -> Left err
  where
    -- This is a helper function to assist us in going from a Lazy ByteString, to a Strict Text
    lazyByteStringToStrictText =
      decodeUtf8 . LBS.toStrict

-- This has a number of benefits, we're able to isolate our validation
-- requirements into smaller components that are simpler to maintain and verify.
-- It also allows for greater reuse and it also means that validation is not
-- duplicated across the application, maybe incorrectly.
mkViewRequest
  :: Text
  -> Either Error RqType
mkViewRequest topicTxt =
  case mkTopic topicTxt of
    Right topic -> Right $ ViewRq topic
    Left err -> Left err

mkListRequest
  :: Either Error RqType
mkListRequest = Right ListRq

mkErrorResponse
  :: Error
  -> Response
mkErrorResponse EmptyTopicError = resp400 PlainTextContent "empty topic"
mkErrorResponse EmptyCommentError = resp400 PlainTextContent "empty comment"
mkErrorResponse InvalidReqError = resp400 PlainTextContent "invalid request"

-- Use our ``RqType`` helpers to write a function that will take the input
-- ``Request`` from the Wai library and turn it into something our application
-- cares about.
mkRequest
  :: Request
  -> IO ( Either Error RqType )
mkRequest req = case (requestMethod req, pathInfo req) of
  ("POST", [topic, "add"]) -> do
    comment <- strictRequestBody req
    return $ mkAddRequest topic comment
  ("GET", [topic, "view"]) ->
    return $ mkViewRequest topic
  ("GET", ["list"]) ->
    return $ mkListRequest
  _ -> return $ Left InvalidReqError



-- If we find that we need more information to handle a request, or we have a
-- new type of request that we'd like to handle then we update the ``RqType``
-- structure and the compiler will let us know which parts of our application
-- are affected.
--
-- Reduction of concerns such that each section of the application only deals
-- with a small piece is one of the benefits of developing in this way.
--
-- For now, return a made-up value for each of the responses as we don't have
-- any persistent storage. Plain text responses that contain "X not implemented
-- yet" should be sufficient.
handleRequest
  :: RqType
  -> Either Error Response
handleRequest (AddRq _ _) = Right $ resp200 PlainTextContent "add not implmented yet"
handleRequest (ViewRq _) = Right $ resp200 PlainTextContent "view not implmented yet"
handleRequest ListRq = Right $ resp200 PlainTextContent "list not implmented yet"

-- Reimplement this function using the new functions and ``RqType`` constructors
-- as a guide.
app
  :: Application
app req cb = do
  res <- mkRequest req
  case res of
    Left err -> cb $ mkErrorResponse err
    Right req' ->
      case handleRequest req' of
        Right resp -> cb resp
        Left err -> cb $ mkErrorResponse err

runApp :: IO ()
runApp = run 3000 app
