{-# LANGUAGE NoImplicitPrelude #-}
{-# LANGUAGE ScopedTypeVariables #-}
{-# LANGUAGE OverloadedStrings #-}
{-# LANGUAGE RebindableSyntax #-}

module Course.FileIO where

import Course.Core
import Course.Applicative
import Course.Monad
import Course.Functor
import Course.List

{-

Useful Functions --

  getArgs :: IO (List Chars)
  putStrLn :: Chars -> IO ()
  readFile :: Chars -> IO Chars
  lines :: Chars -> List Chars
  void :: IO a -> IO ()

Abstractions --
  Applicative, Monad:

    <$>, <*>, >>=, =<<, pure

Problem --
  Given a single argument of a file name, read that file,
  each line of that file contains the name of another file,
  read the referenced file and print out its name and contents.

Example --
Given file files.txt, containing:
  a.txt
  b.txt
  c.txt

And a.txt, containing:
  the contents of a

And b.txt, containing:
  the contents of b

And c.txt, containing:
  the contents of c

$ runhaskell FileIO.hs "files.txt"
============ a.txt
the contents of a

============ b.txt
the contents of b

============ c.txt
the contents of c

-}

-- /Tip:/ use @getArgs@ and @run@
main ::
  IO ()
main = do
  args <- getArgs
  case args of
    (arg :. _) -> run arg
    _ -> error "missing argument"

type FilePath =
  Chars

-- /Tip:/ Use @getFiles@ and @printFiles@.
run ::
  Chars
  -> IO ()
run filesListPath = do
  text <- readFile filesListPath
  let filePaths = lines text
  files <- getFiles filePaths
  printFiles files

getFiles ::
  List FilePath
  -> IO (List (FilePath, Chars))
getFiles Nil = return Nil
getFiles (x :. xs) = do
  file <- getFile x
  files <- getFiles xs
  return $ file :. files


getFile ::
  FilePath
  -> IO (FilePath, Chars)
getFile path = do
  text <- readFile path
  return (path, text)

printFiles ::
  List (FilePath, Chars)
  -> IO ()
printFiles Nil = return ()
printFiles (x :. xs) = do
  printFile (fst x) (snd x)
  printFiles xs

printFile ::
  FilePath
  -> Chars
  -> IO ()
printFile path text = do
  putStrLn $ "============ " ++ path
  putStrLn text
