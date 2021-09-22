# common/keyword_replace.py

#def replaceKeyword(text, lefttoken, righttoken, keyword, replacement):
#    keywordMapper = KeywordMapper('', lefttoken, righttoken)
#    return keywordMapper.replace(text, KeywordReplaceHandler({lefttoken: righttoken}))
    
#def replaceToken(text, lefttoken, righttoken, newlefttoken, newrighttoken):
#    keywordMapper = KeywordMapper('', lefttoken, righttoken)
#    return keywordMapper.replace(text, TokenReplaceHandler(newlefttoken, newrighttoken)

class KeywordMapper:
    def __init__(self, template, lefttag, righttag):
        self.template = template
        self.keywordOnly = False
        self.slefttag = lefttag
        self.srighttag = righttag
        self.slefttaglen = len(lefttag)
        self.srighttaglen = len(righttag)

    def replace(self, template, handler):
        if handler is None:
            return template
        if '' == template :
            template = self.template
            if '' == template :
                return ''
#        print('[KeywordMapper.replace] template = {0}'.format(template))
        self.document = ''
        self.keyword = ''
        scope = False
        boundary = False
        skeyword = ''
        svalue = ''
        handler.startToMap()
        for i in range(0, len(template)):
#            print('template[{0}] = {1}'.format(i, template[i]))
            # maybe find a keyword beginning Tag - now isn't in a keyword
            if True != scope and template[i] == self.slefttag[0] :
                if self.isATag(i, self.slefttag, template) :
                    self.keyword = ''
                    scope = True
            elif scope and template[i] == self.srighttag[0] :
                # maybe find a keyword ending Tag - now in a keyword
                if self.isATag(i, self.srighttag, template) :
                    i = i + self.srighttaglen - 1
                    skeyword = self.keyword[self.slefttaglen:len(self.keyword)]
                    svalue = handler.replace(skeyword)
                    if '' == svalue :
                        svalue = '{0}{1}{2}'.format(self.slefttag, skeyword, self.srighttag)
#                    print('value ->{0}'.format(svalue))
                    self.document += svalue
                    boundary = True
                    scope = False

            if True != boundary :
                if True != scope and True != self.keywordOnly :
                    self.document += template[i]
                else :
                    self.keyword += template[i]
            else :
                boundary = False

#            print('scope = ', scope, ', boundary = ', boundary, ', self.keyword = ', self.keyword)#, ', document = ', self.document)
            if i == len(template)-1 :
                if True == scope :
                    self.document += self.keyword
        handler.endOfMapping(self.document)
        return handler.getResult()
    
    def isATag(self, i, tag, template):
        if len(template) >= len(tag):
            for j in range(0, len(tag)):
                if tag[j] != template[i+j]:
                    return False
            return True
        return False

class KeywordReplaceHandler:
    def __init__(self, keywordMap):
	    self.keywordMap = keywordMap

    def startToMap(self):
	    self.result = ''

    def replace(self, keyword):
#        print('Got keyword =========>{0}'.format(keyword))
        if self.keywordMap[keyword] is not None:
            return self.keywordMap[keyword]
        return ''

    def endOfMapping(self, document):
	    self.result = document

    def getResult(self):
	    return self.result

class TokenReplaceHandler:
    def __init__(self, lefttoken, righttoken):
	    self.lefttoken = lefttoken
	    self.righttoken = righttoken

    def startToMap(self):
	    self.result = ''

    def replace(self, keyword):
        if (self.lefttoken is not None) and (self.righttoken is not None):
            return self.lefttoken+keyword+self.righttoken
        return ''

    def endOfMapping(self, document):
	    self.result = document

    def getResult(self):
	    return self.result
